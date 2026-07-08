document.addEventListener("DOMContentLoaded", () => {
    const copyBtn = document.getElementById("copy-btn");
    const form = document.querySelector("footer form");
    const userTextInput = document.getElementById("user-text");

    if (copyBtn && form && userTextInput) {
        copyBtn.addEventListener("click", async (e) => {
            // ── Prevent default behavior ──────────────────────────────────────────────────⊃
            e.preventDefault();

            // ── Client-side validation check for required empty input text ────────────────⊃
            if (!userTextInput.value.trim()) {
                // Trigger the browser's native operational warning bubble
                userTextInput.reportValidity();
                return;
            }

            // ── Gather all current form elements dynamically into standard POST payload ───⊃
            const formData = new FormData(form);

            try {
                // ── Dispatch background async POST request to the custom copy endpoint ────────⊃
                const response = await fetch("/api/copy", {
                    method: "POST",
                    body: formData
                });

                if (!response.ok) {
                    throw new Error("Server rejected or failed processing the ASCII copy operation");
                }

                // ── Extract the raw text representing f.SimpleResult from response body ────────⊃
                const textToCopy = await response.text();

                // ── Write the unformatted clean ASCII representation straight to clipboard ────⊃
                await navigator.clipboard.writeText(textToCopy);

                // ── Provide temporary visual indicator feedback directly on the button UI ─────⊃
                const originalText = copyBtn.innerText;
                copyBtn.innerText = "DONE";
                copyBtn.style.borderColor = "#2ecc71";
                copyBtn.style.color = "#2ecc71";

                setTimeout(() => {
                    copyBtn.innerText = "COPY";
                    copyBtn.style.borderColor = "";
                    copyBtn.style.color = "";
                }, 900);

            } catch (err) {
                console.error("Async copy sequence aborted:", err);
                alert("Failed to securely retrieve clean layout from the server.");
            }
        });
    }
});
