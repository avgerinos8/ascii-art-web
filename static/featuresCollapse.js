document.addEventListener("DOMContentLoaded", () => {
  // Helper function to handle toggle logic cleanly using strict if conditions
  const setupToggle = (triggerId, targetSelector) => {
    const trigger = document.getElementById(triggerId);

    // Check if the trigger element exists before proceeding
    if (trigger) {
      const target = document.querySelector(targetSelector);

      // Check if the target element exists before proceeding
      if (target) {

        // Function to update directional arrow symbols based on current layout state
        const updateArrows = () => {
          const arrows = trigger.querySelectorAll(".arrow");
          const isCollapsed = target.classList.contains("row-collapsed");

          arrows.forEach(arrow => {
            if (isCollapsed) {
              arrow.textContent = "\u2303"; // Upward arrow symbol
            } else {
              arrow.textContent = "\u2304"; // Downward arrow symbol
            }
          });
        };

        // 1. Evaluate and apply correct arrow directions on initial page layout load
        updateArrows();

        // 2. Listen for manual click interactions to collapse or expand rows
        trigger.addEventListener("click", (e) => {
          e.stopPropagation();
          target.classList.toggle("row-collapsed");
          updateArrows(); // Refresh arrow states immediately after structural transition
        });
      }
    }
  };

  // Bind triggers to their respective elements
  setupToggle("font-toggle", ".optionsLine1");
  setupToggle("options-toggle", ".optionsLine2");
  setupToggle("text-toggle", ".inputTextLine");
  setupToggle("colors-toggle", ".addColor");
});

document.addEventListener("DOMContentLoaded", () => {
  const userText = document.getElementById("user-text");
  const increaseBtn = document.getElementById("row-increase-btn");
  const decreaseBtn = document.getElementById("row-decrease-btn");

  // Check if all necessary interactive elements exist before binding events
  if (userText && increaseBtn && decreaseBtn) {

    // Handle row increment action up to a maximum limit of 6
    increaseBtn.addEventListener("click", () => {
      let currentRows = parseInt(userText.getAttribute("rows")) || 3;
      if (currentRows < 6) {
        userText.setAttribute("rows", currentRows + 1);
      }
    });

    // Handle row decrement action down to a minimum limit of 1
    decreaseBtn.addEventListener("click", () => {
      let currentRows = parseInt(userText.getAttribute("rows")) || 3;
      if (currentRows > 1) {
        userText.setAttribute("rows", currentRows - 1);
      }
    });
  }
});