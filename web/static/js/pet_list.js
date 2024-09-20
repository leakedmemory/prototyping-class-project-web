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

// update pet count on creation and deletion
document.body.addEventListener("htmx:afterOnLoad", (event) => {
  const petCount = event.detail.xhr.getResponseHeader("X-Pet-Count");
  if (petCount !== null) {
    const petCountElement = document.querySelector(".pet-num");
    if (petCountElement) {
      petCountElement.innerHTML = `PETS (${petCount})`;
    }
  }
});
