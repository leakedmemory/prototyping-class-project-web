import { Formik } from "formik";
import { FaLock, FaAt } from "react-icons/fa6";
import { Link } from "react-router-dom";

import "./login.css"
import DefaultButton from "../../components/default_button/DefaultButton"

export default function Login() {
  return (
    <div className="login">
      <h1 className="title-login">LOGIN</h1>
      <div className="inputs-login">
        <Formik
          initialValues={{ email: "", password: "" }}
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
            <form onSubmit={handleSubmit} className="inputs-login">
              <div className="div-input-login">
                <FaAt className="default-icon" />
                <input
                  name="email"
                  className="input-login"
                  onChange={handleChange}
                  onBlur={handleBlur}
                  value={values.email}
                  placeholder="Email"
                />
              </div>
              {errors.email && touched.email && errors.email}
              <div className="div-input-login">
                <FaLock className="default-icon" />
                <input
                  type="password"
                  name="password"
                  className="input-login"
                  onChange={handleChange}
                  onBlur={handleBlur}
                  value={values.password}
                  placeholder="Senha"
                />
              </div>
              {errors.password && touched.password && errors.password}
              <DefaultButton title="Entrar" type="submit" marginBottom={22} />
              <div className="signup-message-login">
                <p>Ainda não possui conta?</p>
                <Link to="/signup" className="default-link">Cadastre-se</Link>
              </div>
            </form>
          )}
        </Formik>
      </div>
    </div>
  )
}
