import { Formik } from "formik";
import { FaUser } from "react-icons/fa";
import { FaLock } from "react-icons/fa6";

import "./login.css"
import "../../styles/default_icon.css"
import DefaultButton from "../../components/default_button/DefaultButton"

export default function Login() {
  return (
    <div className="login">
      <h1 className="title">LOGIN</h1>
      <div className="inputs">
        <Formik
          initialValues={{ username: "", password: "" }}
          validate={_ => { }}
          onSubmit={(_, { setSubmitting }) => {
            setTimeout(() => {
              setSubmitting(false);
            }, 400);
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
            <form onSubmit={handleSubmit} className="inputs">
              <div className="div-input">
                <FaUser className="default-icon"/>
                <input
                  name="username"
                  className="input"
                  onChange={handleChange}
                  onBlur={handleBlur}
                  value={values.username}
                />
              </div>
              {errors.username && touched.username && errors.username}
              <div className="div-input">
                <FaLock className="default-icon"/>
                <input
                  type="password"
                  name="password"
                  className="input"
                  onChange={handleChange}
                  onBlur={handleBlur}
                  value={values.password}
                />
              </div>
              {errors.password && touched.password && errors.password}
              <DefaultButton title="Entrar" type="submit" />
            </form>
          )}
        </Formik>
      </div>
    </div>
  )
}
