import * as React from "react";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import { HiDotsVertical } from "react-icons/hi";

import "./basicMenu.css"

interface basicMenuProps {
  items: String[],
}

export default function BasicMenu({ items }: basicMenuProps) {
  const [anchorEl, setAnchorEl] = React.useState<null | SVGElement>(null);
  const open = Boolean(anchorEl);

  const handleClick = (event: React.MouseEvent<SVGElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <div>
      <HiDotsVertical
        id="dots-basic-menu"
        className="dots-icon-basic-menu"
        aria-controls={open ? "basic-menu" : undefined}
        aria-haspopup="true"
        aria-expanded={open ? "true" : undefined}
        onClick={handleClick}
      />
      <Menu
        id="basic-menu"
        anchorEl={anchorEl}
        open={open}
        onClose={handleClose}
        MenuListProps={{
          "aria-labelledby": "dots-basic-menu",
        }}
      >
        {items.map((item, idx) => {
          if (item.toLowerCase() === "deletar") {
            return (
              <MenuItem
                key={idx}
                onClick={handleClose}
                className="delete-item-basic-menu"
              >{item}</MenuItem>
            )
          }

          return (
            <MenuItem key={idx} onClick={handleClose}>{item}</MenuItem>
          )
        })}
      </Menu>
    </div>
  );
}
