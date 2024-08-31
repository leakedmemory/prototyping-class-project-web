document.addEventListener("DOMContentLoaded", () => {
  const addPetModal = document.getElementById("add-pet-modal");

  document.body.addEventListener("click", (event) => {
    if (event.target.id === "show-pet-modal") {
      addPetModal.showModal();
    }
  });

  const addPetButton = document.getElementById("add-pet");
  const cancelButton = document.getElementById("cancel");

  addPetButton.addEventListener("click", () => {
    addPetModal.close();
  });

  cancelButton.addEventListener("click", () => {
    addPetModal.close();
  });
});
