function initScrollSyncAndLayout() {
    const topWrapper = document.querySelector('.topScrollWrapper');
    const topFakeContent = document.querySelector('.topScrollFakeContent');
    const asciiWrapper = document.querySelector('.asciiWrapper');
    const asciiOutput = document.querySelector('.asciiOutput');
    const footer = document.querySelector('footer');

    if (!topWrapper || !asciiWrapper || !asciiOutput) return;

    // 1. Συγχρονισμός του πλάτους για την πάνω μπάρα scroll
    function syncWidth() {
        topFakeContent.style.width = asciiOutput.scrollWidth + 'px';
    }

    syncWidth();
    window.addEventListener('resize', syncWidth);

    const sizeInput = document.querySelector('input[name="BgSize"]');
    if (sizeInput) sizeInput.addEventListener('input', syncWidth);

    // 2. Live Υπολογισμός του ύψους του Footer και inject στο CSS
    // 2. Live Υπολογισμός του ύψους του Footer και inject στο CSS
    if (footer) {
        // Συνάρτηση που υπολογίζει το πραγματικό ορατό ύψος του footer στην οθόνη
        function updateFooterHeight() {
            // Παίρνουμε τη θέση του footer σε σχέση με το viewport
            const rect = footer.getBoundingClientRect();
            // Το ορατό ύψος είναι το ύψος της οθόνης μείον το σημείο που ξεκινάει το footer
            const visibleHeight = window.innerHeight - rect.top;

            // Ενημερώνουμε τη CSS variable live
            document.documentElement.style.setProperty('--footer-height', visibleHeight + 'px');
        }

        // Watch για αλλαγές στο μέγεθος του footer
        const footerObserver = new ResizeObserver(() => {
            updateFooterHeight();
        });
        footerObserver.observe(footer);

        // Επειδή το CSS animation (transition: transform 1.2s) παίρνει χρόνο να ολοκληρωθεί,
        // τρέχουμε ένα loop κατά τη διάρκεια του animation για να μεγαλώνει ο χώρος ομαλά live!
        window.addEventListener('click', () => {
            // Μόλις ο χρήστης κάνει κλικ (π.χ. στο toggle), τρέχουμε το update για τα επόμενα 1.2 δευτερόλεπτα
            let startTime = Date.now();
            function animateHeight() {
                updateFooterHeight();
                if (Date.now() - startTime < 1300) { // 1300ms για να καλύψει το transition των 1.2s
                    requestAnimationFrame(animateHeight);
                }
            }
            requestAnimationFrame(animateHeight);
        });

        // Αρχικός υπολογισμός στο load
        updateFooterHeight();
    }


    // 3. Αμφίδρομος συγχρονισμός οριζόντιου Scroll
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

// Εκκίνηση layout
window.addEventListener('load', initScrollSyncAndLayout);
