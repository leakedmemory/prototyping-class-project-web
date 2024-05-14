import { Formik } from "formik";
import { FaUser } from "react-icons/fa";
import { FaLock, FaAt } from "react-icons/fa6";
import { Link } from "react-router-dom";

import "./signUp.css";
import DefaultButton from "../../components/default_button/DefaultButton";

export default function SignUp() {
  return (
    <div className="sign-up">
      <h1 className="title-signup">CADASTRO</h1>
      <div className="inputs-signup">
        <Formik
          initialValues={{
            name: "",
            email: "",
            password: "",
            confirm_password: "",
          }}
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
            <form onSubmit={handleSubmit} className="inputs-signup">
              <div className="div-input-signup">
                <FaUser className="default-icon" />
                <input
                  name="name"
                  className="input-signup"
                  onChange={handleChange}
                  onBlur={handleBlur}
                  value={values.name}
                  placeholder="Nome"
                />
              </div>
              {errors.name && touched.name && errors.name}
              <div className="div-input-signup">
                <FaAt className="default-icon" />
                <input
                  name="email"
                  className="input-signup"
                  onChange={handleChange}
                  onBlur={handleBlur}
                  value={values.email}
                  placeholder="Email"
                />
              </div>
              {errors.email && touched.email && errors.email}
              <div className="div-input-signup">
                <FaLock className="default-icon" />
                <input
                  type="password"
                  name="password"
                  className="input-signup"
                  onChange={handleChange}
                  onBlur={handleBlur}
                  value={values.password}
                  placeholder="Senha"
                />
              </div>
              {errors.password && touched.password && errors.password}
              <div className="div-input-signup">
                <FaLock className="default-icon" />
                <input
                  type="password"
                  name="confirm_password"
                  className="input-signup"
                  onChange={handleChange}
                  onBlur={handleBlur}
                  value={values.confirm_password}
                  placeholder="Confirmar Senha"
                />
              </div>
              {errors.confirm_password &&
                touched.confirm_password &&
                errors.confirm_password}
              <DefaultButton
                title="Cadastrar"
                type="submit"
                marginBottom={22}
              />
              <div className="login-message-signup">
                <p>Já possui conta?</p>
                <Link to="/login" className="default-link">
                  Fazer Login
                </Link>
              </div>
            </form>
          )}
        </Formik>
      </div>
    </div>
  );
}
