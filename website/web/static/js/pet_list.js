document.addEventListener("DOMContentLoaded", () => {
  const addPetModal = document.getElementById("add-pet-modal");
  const deletePetModal = document.getElementById("delete-pet-modal");

  document.body.addEventListener("click", (event) => {
    if (event.target.id === "show-add-pet-modal") {
      addPetModal.showModal();
    }
  });

  const addPetConfirmButton = document.getElementById("add-pet-confirm");
  const addPetCancelButton = document.getElementById("add-pet-cancel");

  addPetConfirmButton.addEventListener("click", () => {
    addPetModal.close();
  });

  addPetCancelButton.addEventListener("click", () => {
    addPetModal.close();
  });
});
