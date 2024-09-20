document.addEventListener("htmx:responseError", () => {
  const error = document.getElementById("signup-error-message");
  error.style.display = "block";
  error.style.visibility = "visible";
});
