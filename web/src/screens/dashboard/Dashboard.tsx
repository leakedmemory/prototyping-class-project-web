import { PiUserCircleFill } from "react-icons/pi";
import { MdOutlineAddCircle } from "react-icons/md";
import { useState } from "react";
import { Formik } from "formik";

import "./dashboard.css";
import PetCard, {
  petCardDataProps,
  petCardProps,
} from "../../components/pet_card/PetCard";
import BaseModal from "../../components/modal/BaseModal";
import DefaultButton from "../../components/default_button/DefaultButton";

export default function Dashboard() {
  const [userProfile] = useState({
    name: "Dudu",
    email: "duduzinho@exemplo.com",
    number: "(83) 90000-0001",
  });

  const [pets, setPets] = useState<petCardProps[]>([
    {
      name: "Sagwa",
      animalType: "Gato",
      age: "1",
      breed: "Siamês",
    },
    {
      name: "Sagwa",
      animalType: "Gato",
      age: "2",
      breed: "Siamês",
    },
    {
      name: "Sagwa",
      animalType: "Gato",
      age: "3",
      breed: "Siamês",
    },
    {
      name: "Sagwa",
      animalType: "Gato",
      age: "4",
      breed: "Siamês",
    },
    {
      name: "Sagwa",
      animalType: "Gato",
      age: "5",
      breed: "Siamês",
    },
  ]);

  const [modalAddPet, toggleModalAddPet] = useState(false);

  return (
    <div className="dashboard" id="root-dashboard">
      <div className="profile-dashboard">
        <PiUserCircleFill style={{ width: 100, height: 100 }} />
        <div className="profile-data-dashboard">
          <p>{userProfile.name}</p>
          <p>{userProfile.email}</p>
          <p>{userProfile.number}</p>
        </div>
      </div>
      <div className="pets-header-dashboard">
        <p style={{ fontSize: 20, fontWeight: "bold", lineHeight: 0 }}>
          PETS ({pets.length})
        </p>
        <MdOutlineAddCircle
          className="add-pet-dashboard"
          onClick={() => toggleModalAddPet(true)}
        />
      </div>
      <div>
        {pets
          .map((data, idx) => {
            return (
              <PetCard
                key={idx}
                editPet={(data) =>
                  setPets([...pets.slice(0, idx), data, ...pets.slice(idx + 1)])
                }
                deletePet={() => {
                  setPets([...pets.slice(0, idx), ...pets.slice(idx + 1)])
                }}
                {...data}
              />
            );
          })
          .reverse()}
      </div>
      <BaseModal
        open={modalAddPet}
        toggleModal={(isOpen: boolean) => toggleModalAddPet(isOpen)}
      >
        <h2>Adicionar pet</h2>
        <Formik
          initialValues={{ name: "", breed: "", age: "", animalType: "" }}
          onSubmit={(data: petCardDataProps) => {
            setPets([...pets, data]);
            toggleModalAddPet(false);
          }}
        >
          {({
            values,
            errors,
            touched,
            handleChange,
            handleBlur,
            handleSubmit,
          }) => (
            <form onSubmit={handleSubmit} className="inputs-login">
              <input
                name="name"
                className="input-edit-pet"
                onChange={handleChange}
                onBlur={handleBlur}
                value={values.name}
                placeholder="Nome"
              />
              {errors.name && touched.name && errors.name}
              <input
                name="breed"
                className="input-edit-pet"
                onChange={handleChange}
                onBlur={handleBlur}
                value={values.breed}
                placeholder="Raça"
              />
              {errors.breed && touched.breed && errors.breed}
              <input
                name="age"
                className="input-edit-pet"
                onChange={handleChange}
                onBlur={handleBlur}
                value={values.age}
                placeholder="Idade"
              />
              {errors.age && touched.age && errors.age}
              <input
                name="animalType"
                className="input-edit-pet"
                onChange={handleChange}
                onBlur={handleBlur}
                value={values.animalType}
                placeholder="Tipo"
              />
              {errors.animalType && touched.animalType && errors.animalType}
              <DefaultButton title="Criar" type="submit" />
            </form>
          )}
        </Formik>
      </BaseModal>
    </div>
  );
}
