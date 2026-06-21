function initFontWrapControl() {
    const wrapToggle = document.getElementById('font-wrap-toggle');
    const alignSelect = document.getElementById('font-align-select');

    // Έλεγχος αν υπάρχουν τα στοιχεία στο DOM για να μην κρασάρει το script
    if (!wrapToggle || !alignSelect) return;

    // Συνάρτηση που ελέγχει την κατάσταση του checkbox και ρυθμίζει το dropdown
    function updateSelectState() {
        // Αν το checkbox είναι checked, το disabled είναι false (ενεργό)
        // Αν το checkbox ΔΕΝ είναι checked, το disabled είναι true (grayed out)
        alignSelect.disabled = !wrapToggle.checked;
    }

    // Εκτέλεση κατά το πρώτο φόρτωμα της σελίδας για να συγχρονιστούν σωστά
    updateSelectState();

    // Ακρόαση (Listener) για κάθε φορά που ο χρήστης κάνει κλικ στο Font Wrap
    wrapToggle.addEventListener('change', updateSelectState);
}

// Εκκίνηση της συνάρτησης μόλις φορτώσει όλη η σελίδα
window.addEventListener('load', initFontWrapControl);
