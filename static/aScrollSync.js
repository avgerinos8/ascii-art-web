function initScrollSyncAndLayout() {
    const topWrapper = document.querySelector('.topScrollWrapper');
    const topFakeContent = document.querySelector('.topScrollFakeContent');
    const asciiWrapper = document.querySelector('.asciiWrapper');
    const asciiOutput = document.querySelector('.asciiOutput');
    const footer = document.querySelector('footer');

    if (!topWrapper || !asciiWrapper || !asciiOutput) return;

    // 1. Synchronize fake top scrollbar content width with actual layout width
    function syncWidth() {
        // Run inside requestAnimationFrame to align execution with the next browser layout repaint
        requestAnimationFrame(() => {
            if (topFakeContent && asciiOutput) {
                const actualWidth = asciiOutput.scrollWidth;
                topFakeContent.style.width = actualWidth + 'px';
            }
        });
    }

    // 2. Track background fetch operations using asynchronous execution loops
    // MutationObserver catches direct manipulations onto innerHTML or textContent
    const contentObserver = new MutationObserver(() => {
        syncWidth();
        // Double-check 50ms later in case custom typography rendering updates layout bounds slowly
        setTimeout(syncWidth, 50);
    });
    contentObserver.observe(asciiOutput, { childList: true, characterData: true, subtree: true });

    // ResizeObserver tracks explicit shifts on container layout metrics
    const resizeObserver = new ResizeObserver(() => {
        syncWidth();
    });
    resizeObserver.observe(asciiOutput);

    // Force evaluation execution cycles to wait until system or network typography formats load completely
    if (document.fonts) {
        document.fonts.ready.then(() => {
            syncWidth();
            setTimeout(syncWidth, 100);
        });
    } else {
        syncWidth();
    }

    window.addEventListener('resize', syncWidth);

    const sizeInput = document.querySelector('input[name="BgSize"]');
    if (sizeInput) sizeInput.addEventListener('input', syncWidth);

    // 3. Live calculations rendering visible footer dimensional structures directly into CSS DOM targets
    if (footer) {
        function updateFooterHeight() {
            const rect = footer.getBoundingClientRect();
            const visibleHeight = window.innerHeight - rect.top;
            document.documentElement.style.setProperty('--footer-height', visibleHeight + 'px');
        }

        const footerObserver = new ResizeObserver(() => {
            updateFooterHeight();
        });
        footerObserver.observe(footer);

        window.addEventListener('click', () => {
            let startTime = Date.now();
            function animateHeight() {
                updateFooterHeight();
                if (Date.now() - startTime < 1300) {
                    requestAnimationFrame(animateHeight);
                }
            }
            requestAnimationFrame(animateHeight);
        });

        updateFooterHeight();
    }

    // 4. Two-way operational tracking loop handling global horizontal page scrolls
    let isSyncingTop = false;
    let isSyncingAscii = false;

    topWrapper.addEventListener('scroll', () => {
        if (!isSyncingTop) {
            isSyncingAscii = true;
            asciiWrapper.scrollLeft = topWrapper.scrollLeft;
        }
        isSyncingTop = false;
    });

    asciiWrapper.addEventListener('scroll', () => {
        if (!isSyncingAscii) {
            isSyncingTop = true;
            topWrapper.scrollLeft = asciiWrapper.scrollLeft;
        }
        isSyncingAscii = false;
    });
}

// Global invocation sequence triggered automatically when asset compilation loops finish loading
window.addEventListener('load', initScrollSyncAndLayout);
