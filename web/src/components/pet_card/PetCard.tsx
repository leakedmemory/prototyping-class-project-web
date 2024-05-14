import { PiUserCircleFill } from "react-icons/pi";
import { useState } from "react";
import { Formik } from "formik";

import "./petCard.css";
import editIcon from "../../assets/editIcon.svg";
import deleteIcon from "../../assets/deleteIcon.svg";
import BaseModal from "../../components/modal/BaseModal";
import DefaultButton from "../default_button/DefaultButton";

export interface petCardDataProps {
  name: string;
  animalType: string;
  age: string;
  breed: string;
}

export interface petCardProps extends petCardDataProps {
  editPet?: (petCardDataProps: petCardDataProps) => void;
}

export default function PetCard({
  name,
  animalType,
  age,
  breed,
  editPet,
}: petCardProps) {
  const [modalEditPet, toggleModalEditPet] = useState(false);

  return (
    <div className="petcard" id="root-petcard">
      <PiUserCircleFill style={{ width: 60, height: 60 }} />
      <div className="profile-petcard">
        <div className="profile-data-petcard">
          <div className="profile-column-petcard">
            <p>{name}</p>
            <p>{animalType}</p>
          </div>
          <div className="profile-column-petcard">
            <p>{age} ano(s)</p>
            <p>{breed}</p>
          </div>
        </div>
        <div className="icon-div-petcard">
          <img
            src={editIcon}
            alt="edit icon"
            className="icon-petcard"
            onClick={() => toggleModalEditPet(true)}
          />
          <img src={deleteIcon} alt="delete icon" className="icon-petcard" />
        </div>
      </div>
      <BaseModal
        open={modalEditPet}
        toggleModal={(isOpen: boolean) => toggleModalEditPet(isOpen)}
      >
        <h2>Editar pet</h2>
        <Formik
          initialValues={{ name, breed, age, animalType }}
          onSubmit={(data: petCardDataProps) => {
            editPet!(data);
            toggleModalEditPet(false);
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
              <DefaultButton title="Editar" type="submit" />
            </form>
          )}
        </Formik>
      </BaseModal>
    </div>
  );
}
