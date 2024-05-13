import { PiUserCircleFill } from "react-icons/pi";

import "./petCard.css"
import BasicMenu from "../basic_menu/BasicMenu"

export interface petCardProps {
  name: String,
  type: String,
  age: number,
  breed: String,
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
        <BasicMenu items={["Editar", "Deletar"]} />
      </div>
    </div>
  )
}
