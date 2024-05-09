import React from "react"
import ReactDOM from "react-dom/client"
import Login from "./screens/login/Login"
import SignUp from "./screens/sign_up/SignUp"
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import "./index.css"

const router = createBrowserRouter([
  {
    path: "/",
    element: <Login />,
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/signup",
    element: <SignUp />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
