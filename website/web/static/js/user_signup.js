document.addEventListener("htmx:responseError", () => {
  const error = document.getElementById("error-message");
  error.style.display = "block";
  error.style.visibility = "visible";
});
