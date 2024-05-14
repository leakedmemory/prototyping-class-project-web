import { PiUserCircleFill } from "react-icons/pi";
import { MdOutlineAddCircle } from "react-icons/md";

import "./dashboard.css";
import PetCard, { petCardProps } from "../../components/pet_card/PetCard";
import { useState } from "react";

const userPets: petCardProps[] = [
  {
    name: "Sagwa",
    type: "Gato",
    age: 3,
    breed: "Siamês",
  },
];

export default function Dashboard() {
  const [userProfile] = useState({
    name: "Dudu",
    email: "duduzinho@exemplo.com",
    number: "(83) 90000-0001",
  });

  return (
    <div className="dashboard">
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
          PETS ({userPets.length})
        </p>
        <MdOutlineAddCircle className="add-pet-dashboard" />
      </div>
      <div>
        {userPets.map((data, idx) => {
          return <PetCard key={idx} {...data} />;
        })}
      </div>
    </div>
  );
}
