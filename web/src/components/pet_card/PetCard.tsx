import { PiUserCircleFill } from "react-icons/pi";

import "./petCard.css";

import editIcon from "../../assets/editIcon.svg";
import deleteIcon from "../../assets/deleteIcon.svg";

export interface petCardProps {
  name: string;
  type: string;
  age: number;
  breed: string;
}

export default function PetCard({ name, type, age, breed }: petCardProps) {
  return (
    <div className="petcard">
      <PiUserCircleFill style={{ width: 60, height: 60 }} />
      <div className="profile-petcard">
        <div className="profile-data-petcard">
          <div className="profile-column-petcard">
            <p>{name}</p>
            <p>{type}</p>
          </div>
          <div className="profile-column-petcard">
            <p>{age} ano(s)</p>
            <p>{breed}</p>
          </div>
        </div>
        <div className="icon-div-petcard">
          <img src={editIcon} alt="edit icon" className="icon-petcard" />
          <img src={deleteIcon} alt="delete icon" className="icon-petcard" />
        </div>
      </div>
    </div>
  );
}
