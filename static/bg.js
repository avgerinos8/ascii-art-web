// ── main ────────────────────────────────────────────────────────────────────⊃
// ── Background ASCII Vignette (Footer Only) ─────────────────────────────────⊃

// 1. TWEAK THESE VALUES TO CHANGE THE EFFECT
const CONFIG = {
    bgSize: 1.5,         // Adjusted scale of the ASCII characters
    bgDistance: 1,       // Adjusted vertical line spacing
    noise: 0.30,         // Static spatial variance / grain
    breatheSpeed: 0.0001, // Adjusted speed of the wave animation
    breatheDepth: 0.45,  // Intensity of the wave breathing effect
    waveSpread: 0.6,     // Distance between wave peaks
    mouseRadius: 12,     // Size of the mouse influence area in grid cells
    mousePeak: 0.4,      // Maximum brightness added by the mouse cursor
    mouseDecay: 0.001,   // How quickly the mouse thermal trail fades per frame
    softness: 7.5        // Gaussian falloff sharpness (higher = softer edge)
};

// ── Internal Derived Variables ────────⊃
const CHARS = [' ', ' ', '░', '░', '▒', '▓', '█'];
const OP = 0.35; // Base opacity
const FS = 16 * CONFIG.bgSize;
const LINE_HEIGHT = FS * CONFIG.bgDistance;
const CW = FS * 0.6;

// Helper function to convert HEX color string to RGB object
function hexToRgb(hex) {
    const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex.trim());
    return result ? {
        r: parseInt(result[1], 16),
        g: parseInt(result[2], 16),
        b: parseInt(result[3], 16)
    } : { r: 145, g: 145, b: 145 }; // Default fallback if parsing fails
}

// Read the initial hex value directly from the new CSS variable
let currentHex = getComputedStyle(document.documentElement).getPropertyValue('--panel-effect-color');
let rgbColor = hexToRgb(currentHex);

// Listen to the custom event triggered by the theme toggle script
// Fetch the updated variable directly from the DOM to ensure we get panel-effect-color
window.addEventListener('themeChanged', () => {
    const updatedHex = getComputedStyle(document.documentElement).getPropertyValue('--panel-effect-color');
    rgbColor = hexToRgb(updatedHex);
});

// Inject canvas specifically into the footer
const footer = document.querySelector('footer');
const canvas = document.createElement('canvas');

// Absolute positioning makes it fill the footer precisely
canvas.style.cssText = 'position:absolute;top:0;left:0;width:100%;height:100%;z-index:-1;pointer-events:none;user-select:none;';

// Force footer to clip the canvas at the rounded corners
if (footer) {
    footer.style.overflow = 'hidden';

    // Ensure the form content inside footer sits on top of the canvas
    const formContainer = footer.querySelector('.FooterContainer');
    if (formContainer) {
        formContainer.style.position = 'relative';
        formContainer.style.zIndex = '1';
    }

    footer.appendChild(canvas);
}

const ctx = canvas.getContext('2d');

let noiseGrid = [];
let heatGrid = [];
let cols = 0;
let rows = 0;

function buildGrid() {
    if (!footer) return;

    // Use footer's actual pixel dimensions
    canvas.width = footer.clientWidth;
    canvas.height = footer.clientHeight;

    cols = Math.floor(canvas.width / CW) + 1;
    rows = Math.floor(canvas.height / LINE_HEIGHT) + 1;
    const newNoiseGrid = [];
    const newHeatGrid = [];

    for (let y = 0; y < rows; y++) {
        newNoiseGrid[y] = [];
        newHeatGrid[y] = [];
        for (let x = 0; x < cols; x++) {
            newNoiseGrid[y][x] = {
                spatialNoise: (Math.random() - 0.5) * CONFIG.noise,
                phaseOffset: Math.random() * Math.PI * 2,
            };
            newHeatGrid[y][x] = 0;
        }
    }

    // Atomic swap to prevent animation frame race conditions
    noiseGrid = newNoiseGrid;
    heatGrid = newHeatGrid;
}

// ── On mouse move: inject entropy with soft gaussian brush ────────⊃
window.addEventListener('mousemove', e => {
    if (heatGrid.length === 0 || !heatGrid[0]) return;

    // Get footer's exact position on the screen
    const rect = canvas.getBoundingClientRect();

    // Convert mouse coordinates to be relative to the footer's top-left corner
    const localX = e.clientX - rect.left;
    const localY = e.clientY - rect.top;

    // Convert local px position to grid coords
    const gx = localX / CW;
    const gy = localY / LINE_HEIGHT;

    const igx = Math.floor(gx);
    const igy = Math.floor(gy);

    for (let dy = -CONFIG.mouseRadius; dy <= CONFIG.mouseRadius; dy++) {
        for (let dx = -CONFIG.mouseRadius; dx <= CONFIG.mouseRadius; dx++) {
            const nx = igx + dx;
            const ny = igy + dy;
            if (nx < 0 || ny < 0 || nx >= cols || ny >= rows) continue;

            const screenDx = dx * CW;
            const screenDy = dy * LINE_HEIGHT;
            const screenR = CONFIG.mouseRadius * Math.max(CW, LINE_HEIGHT);

            const d2 = (screenDx * screenDx + screenDy * screenDy) / (screenR * screenR);
            const heat = Math.exp(-d2 * CONFIG.softness) * CONFIG.mousePeak;

            if (heat > heatGrid[ny][nx]) heatGrid[ny][nx] = heat;
        }
    }
});

function draw(timestamp) {
    if (!ctx) return;

    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.font = `${FS}px monospace`;

    const cx = cols / 2;
    const cy = rows / 2;
    const maxD = Math.hypot(cx, cy);

    for (let y = 0; y < rows; y++) {
        for (let x = 0; x < cols; x++) {
            const cell = noiseGrid[y]?.[x];
            if (!cell) continue;

            const dist = Math.hypot(x - cx, y - cy) / maxD;
            const wave = Math.sin(timestamp * CONFIG.breatheSpeed + dist * CONFIG.waveSpread + cell.phaseOffset);

            const heat = heatGrid[y]?.[x] ?? 0;
            const t = Math.max(0, Math.min(1, 0.75 * dist + cell.spatialNoise + wave * CONFIG.breatheDepth + heat));

            const charIndex = Math.floor(t * (CHARS.length - 1));
            const opacity = Math.min(1, OP + wave * 0.08 + heat * 0.2);

            ctx.fillStyle = `rgba(${rgbColor.r}, ${rgbColor.g}, ${rgbColor.b}, ${opacity.toFixed(3)})`;
            ctx.fillText(CHARS[charIndex], x * CW, (y + 1) * LINE_HEIGHT);

            if (heatGrid[y][x] > 0) heatGrid[y][x] = Math.max(0, heatGrid[y][x] - CONFIG.mouseDecay);
        }
    }

    requestAnimationFrame(draw);
}

// ── Initialization & Auto-Resize ────────────────────────────────────────────⊃
function init() {
    if (!footer) return;
    buildGrid();
    requestAnimationFrame(draw);
}

// Observe the footer element specifically to rebuild grid perfectly when it collapses or expands
if (footer) {
    const observer = new ResizeObserver(() => {
        buildGrid();
    });
    observer.observe(footer);
}

init();