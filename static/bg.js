// ── Background ASCII Vignette — Breathing + Mouse Entropy ────────⊃

// Helper function to convert HEX color string (6 or 8 characters) to RGB object
function hexToRgb(hex) {
    // Updated regex to optionally capture the alpha channel (last 2 characters) so it doesn't fail on colors like #7a7a7a2f
    const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})?$/i.exec(hex.trim());
    return result ? {
        r: parseInt(result[1], 16),
        g: parseInt(result[2], 16),
        b: parseInt(result[3], 16)
        // Alpha channel is intentionally ignored here since opacity is calculated dynamically in the draw loop
    } : { r: 145, g: 145, b: 145 }; // Default fallback to #919191 if parsing fails
}

// Helper function to determine the correct CSS variable based on the active theme
function getActiveThemeColor() {
    const isDark = document.body.getAttribute('data-theme') === 'dark';
    const varName = isDark ? '--dark-bg-effe' : '--light-bg-effe';
    return getComputedStyle(document.documentElement).getPropertyValue(varName).trim();
}

// Read the initial hex value directly based on the current theme state
let currentHex = getActiveThemeColor();
let rgbColor = hexToRgb(currentHex);

// Listen to the custom event triggered by the theme toggle script to update color reactively
window.addEventListener('themeChanged', () => {
    // A slight delay ensures the data-theme attribute on the body has been updated before we read it
    setTimeout(() => {
        currentHex = getActiveThemeColor();
        rgbColor = hexToRgb(currentHex);
    }, 0);
});

const CHARS = [' ', ' ', '░', '░', '▒', '▓', '█'];

// Core structural variables controlled by sliders
let BGSIZE = 0.5;
let BgDistance = 2;
let FS = 16 * BGSIZE;
let LINE_HEIGHT = FS * BgDistance;
let CW = FS * 0.6;
const OP = 0.35;
let NOISE = 0.30; // Dynamic noise grain variable

// ── Breathing variables ────────⊃
let BREATHE_SPEED = 0.0004;
const WAVE_SPREAD = 0.6; // Kept as a sweet-spot constant
let BREATHE_DEPTH = 0.25;

// ── Mouse entropy constants ────────⊃
const MOUSE_RADIUS = 12;     // influence radius in grid cells
const MOUSE_PEAK = 0.4;   // max entropy boost at cursor
const MOUSE_DECAY = 0.001; // how fast disturbance fades per frame
const SOFTNESS = 7.5;   // gaussian falloff sharpness — higher = softer edges

const canvas = document.createElement('canvas');
canvas.style.cssText = 'position:fixed;top:0;left:0;z-index:0;pointer-events:none;user-select:none;';
document.body.appendChild(canvas);

const ctx = canvas.getContext('2d');

let noiseGrid = [];
let heatGrid = [];

function buildGrid() {
    const cols = Math.floor(canvas.width / CW) + 1;
    const rows = Math.floor(canvas.height / LINE_HEIGHT) + 1;
    noiseGrid = [];
    heatGrid = [];
    for (let y = 0; y < rows; y++) {
        noiseGrid[y] = [];
        heatGrid[y] = [];
        for (let x = 0; x < cols; x++) {
            noiseGrid[y][x] = {
                spatialNoise: (Math.random() - 0.5) * NOISE,
                phaseOffset: Math.random() * Math.PI * 2,
            };
            heatGrid[y][x] = 0;
        }
    }
}

// ── On mouse move: inject entropy with soft gaussian brush ────────⊃
window.addEventListener('mousemove', e => {
    const cols = heatGrid[0]?.length ?? 0;
    const rows = heatGrid.length;

    // Convert mouse px position to grid coords
    const gx = e.clientX / CW;
    const gy = e.clientY / LINE_HEIGHT;

    const igx = Math.floor(gx);
    const igy = Math.floor(gy);

    for (let dy = -MOUSE_RADIUS; dy <= MOUSE_RADIUS; dy++) {
        for (let dx = -MOUSE_RADIUS; dx <= MOUSE_RADIUS; dx++) {
            const nx = igx + dx;
            const ny = igy + dy;
            if (nx < 0 || ny < 0 || nx >= cols || ny >= rows) continue;

            // Normalize dx/dy by cell aspect ratio to get a true circle in screen space
            const screenDx = dx * CW;
            const screenDy = dy * LINE_HEIGHT;
            const screenR = MOUSE_RADIUS * Math.max(CW, LINE_HEIGHT);

            // Gaussian falloff: e^(-d² * softness) — no hard edge, feathers to 0 smoothly
            const d2 = (screenDx * screenDx + screenDy * screenDy) / (screenR * screenR);
            const heat = Math.exp(-d2 * SOFTNESS) * MOUSE_PEAK;

            if (heat > heatGrid[ny][nx]) heatGrid[ny][nx] = heat;
        }
    }
});

function draw(timestamp) {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    ctx.font = `${FS}px monospace`;

    const cols = Math.floor(canvas.width / CW) + 1;
    const rows = Math.floor(canvas.height / LINE_HEIGHT) + 1;
    const cx = cols / 2;
    const cy = rows / 2;
    const maxD = Math.hypot(cx, cy);

    for (let y = 0; y < rows; y++) {
        for (let x = 0; x < cols; x++) {
            const cell = noiseGrid[y]?.[x];
            if (!cell) continue;

            const dist = Math.hypot(x - cx, y - cy) / maxD;
            const wave = Math.sin(timestamp * BREATHE_SPEED + dist * WAVE_SPREAD + cell.phaseOffset);

            const heat = heatGrid[y]?.[x] ?? 0;
            const t = Math.max(0, Math.min(1, 0.75 * dist + cell.spatialNoise + wave * BREATHE_DEPTH + heat));

            const charIndex = Math.floor(t * (CHARS.length - 1));
            const opacity = Math.min(1, OP + wave * 0.08 + heat * 0.2);

            // Uses the locally saved, reactive rgbColor values for ultra-fast rendering performance
            ctx.fillStyle = `rgba(${rgbColor.r}, ${rgbColor.g}, ${rgbColor.b}, ${opacity.toFixed(3)})`;
            ctx.fillText(CHARS[charIndex], x * CW, (y + 1) * LINE_HEIGHT);

            // Decay heat back to zero
            if (heatGrid[y][x] > 0) heatGrid[y][x] = Math.max(0, heatGrid[y][x] - MOUSE_DECAY);
        }
    }

    requestAnimationFrame(draw);
}

// ── Sliders Input Listeners ─────────────────────────────────────────────────⊃
function setupSliderListeners() {
    const sizeInput = document.querySelector('input[name="BgSize"]');
    const distInput = document.querySelector('input[name="BgDistance"]');
    const speedInput = document.querySelector('input[name="BreatheSpeed"]');
    const noiseInput = document.querySelector('input[name="Noise"]');
    const depthInput = document.querySelector('input[name="BreatheDepth"]');

    if (sizeInput) {
        sizeInput.addEventListener('input', (e) => {
            BGSIZE = parseFloat(e.target.value);
            FS = 16 * BGSIZE;
            CW = FS * 0.6;
            LINE_HEIGHT = FS * BgDistance;
            buildGrid(); // Structure changed, must rebuild matrix mapping
        });
    }

    if (distInput) {
        distInput.addEventListener('input', (e) => {
            BgDistance = parseFloat(e.target.value);
            LINE_HEIGHT = FS * BgDistance;
            buildGrid(); // Line spacing changed, must rebuild matrix mapping
        });
    }

    if (speedInput) {
        speedInput.addEventListener('input', (e) => {
            BREATHE_SPEED = parseFloat(e.target.value);
        });
    }

    if (noiseInput) {
        noiseInput.addEventListener('input', (e) => {
            NOISE = parseFloat(e.target.value);
            buildGrid(); // Re-seed spatial deviation instantly
        });
    }

    if (depthInput) {
        depthInput.addEventListener('input', (e) => {
            BREATHE_DEPTH = parseFloat(e.target.value);
        });
    }
}

function init() {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    // Initialize values from sliders HTML values if they exist
    const sizeInput = document.querySelector('input[name="BgSize"]');
    const distInput = document.querySelector('input[name="BgDistance"]');
    const noiseInput = document.querySelector('input[name="Noise"]');
    const speedInput = document.querySelector('input[name="BreatheSpeed"]');
    const depthInput = document.querySelector('input[name="BreatheDepth"]');

    if (sizeInput) BGSIZE = parseFloat(sizeInput.value);
    if (distInput) BgDistance = parseFloat(distInput.value);
    if (noiseInput) NOISE = parseFloat(noiseInput.value);
    if (speedInput) BREATHE_SPEED = parseFloat(speedInput.value);
    if (depthInput) BREATHE_DEPTH = parseFloat(depthInput.value);

    FS = 16 * BGSIZE;
    LINE_HEIGHT = FS * BgDistance;
    CW = FS * 0.6;

    buildGrid();
    setupSliderListeners();
    requestAnimationFrame(draw);
}

window.addEventListener('resize', () => {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
    buildGrid();
});

init();
