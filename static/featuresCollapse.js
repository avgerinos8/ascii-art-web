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

  // Bind each badge trigger to its respective content row
  setupToggle("font-toggle", ".optionsLine1");
  setupToggle("options-toggle", ".optionsLine2");
  setupToggle("text-toggle", ".inputTextLine");
  setupToggle("colors-toggle", ".addColor");
});
