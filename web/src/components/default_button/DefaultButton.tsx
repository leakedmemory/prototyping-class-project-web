import "./defaultButton.css";

interface props {
  title: string;
  type?: "submit" | "reset" | "button";
  marginBottom?: number;
}

export default function DefaultButton({ title, type, marginBottom }: props) {
  return (
    <>
      <button
        className="default-button"
        type={type}
        style={{ marginBottom: marginBottom ?? 0 }}
      >
        {title}
      </button>
    </>
  );
}
