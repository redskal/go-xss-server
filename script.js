window.onload = function () {
  // Create the overlay container
  const overlay = document.createElement("div");
  overlay.id = "overlay";
  overlay.style.position = "fixed";
  overlay.style.top = "0";
  overlay.style.left = "0";
  overlay.style.width = "40%";
  overlay.style.height = "30%";
  overlay.style.backgroundColor = "rgba(255, 255, 255, 0.95)";
  overlay.style.border = "5px solid #0a0a0a";
  overlay.style.zIndex = "1000";
  overlay.style.display = "flex";
  overlay.style.alignItems = "center";
  overlay.style.justifyContent = "center";
  overlay.style.color = "#000";
  overlay.style.fontFamily = "Segoe UI, Roboto, Helvetica Neue, Ubuntu, sans-serif";
  overlay.style.textAlign = "center";

  // center the overlay
  overlay.style.top = "50%";
  overlay.style.left = "50%";
  overlay.style.transform = "translate(-50%, -50%)";

  // Create the overlay content
  const content = document.createElement("div");
  content.innerHTML = `
        <h1>Remote XSS Payload</h1>
        <p>This XSS payload has been loaded from a remote domain controlled by third parties.</p>
        `;

  // Append content to overlay
  overlay.appendChild(content);

  // Append overlay to body
  document.body.appendChild(overlay);

  // Hide overlay after 10 seconds
  setTimeout(() => {
    overlay.style.display = "none";
  }, 10000);
};
