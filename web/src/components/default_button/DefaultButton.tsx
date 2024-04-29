import "./defaultButton.css"

interface propsButton {
  title: String;
  type: "submit" | "reset" | "button" | undefined;
}

export default function DefaultButton({ title, type }: propsButton) {
  return (
    <>
      <button className="default-button" type={type}>
        {title}
      </button>
    </>
  )
}
