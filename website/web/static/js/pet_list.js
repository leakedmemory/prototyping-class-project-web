// add pet modal
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

// delete pet modal
document.addEventListener("DOMContentLoaded", () => {
  const deletePetModal = document.getElementById("delete-pet-modal");

  document.body.addEventListener("click", (event) => {
    if (event.target.classList.contains("pet-card-delete-icon")) {
      const petName = event.target.getAttribute("petname");
      console.log(petName);
      document.getElementById("delete-pet-modal-petname").textContent = petName;
      deletePetModal.showModal();
    }
  });

  const deletePetConfirmButton = document.getElementById("delete-pet-confirm");
  const deletePetCancelButton = document.getElementById("delete-pet-cancel");

  deletePetConfirmButton.addEventListener("click", () => {
    deletePetModal.close();
  });

  deletePetCancelButton.addEventListener("click", () => {
    deletePetModal.close();
  });
});
