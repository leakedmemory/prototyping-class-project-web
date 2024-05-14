import Box from "@mui/material/Box";
import Modal from "@mui/material/Modal";

const style = {
  borderRadius: 5,
  p: 4,
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  bgcolor: "var(--white)",
  outline: 0,
};

interface BaseModalProps {
  open: boolean;
  toggleModal: (isOpen: boolean) => void;
  // @ts-expect-error: Cannot find 'element'
  children: element;
}

export default function BaseModal(props: BaseModalProps) {
  return (
    <>
      <Modal
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
        open={props.open}
        onClose={() => props.toggleModal(false)}
      >
        <Box sx={{ ...style }}>{props.children}</Box>
      </Modal>
    </>
  );
}
