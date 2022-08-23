import React, { useState } from "react";
import Input from "../../components/input";
import Button from "../../components/Button";
import userService from "../../services/user.service";
import { useRouter } from 'next/router'
import Modal from "../../components/Modal";

const Index = () => {
  const router = useRouter()
  const [user, setUser] = useState({});
  const [showModal, setShowModal] = useState(false);

  const submitRegister = (e) => {
    e.preventDefault();
    userService.register(user)
      .then(
        (data) => {
          // si réponse contient des erreurs, j'affiche ma modal
          if (data.error) {
            setShowModal(true);
          }
          // sinon (si mon utilisateur est inscrit), 
          // je redirige l'utilisateur sur sa page profil et je stocke le jwt pour
          // accéder à ses informations
          else {
            localStorage.setItem('token', data.jwt);
            router.push('/profil')
          }
        }
      )
      .catch(
        //Dans le cas où on aurait des erreurs de type server j'affiche ma modal
        (err) => {
          setShowModal(true);
          console.log(err)
        });
  }
  return (
    <div className="page__register">
      <Modal title="Erreur" isActive={showModal} closeFunction={() => setShowModal(!showModal)} type="information">
        <p>Une erreur est survenue, veuillez contacter le service client.</p>
      </Modal>
      <div className="square">
      </div>
      <div className="square">
      </div>
      <form className="form" onSubmit={(e) => submitRegister(e)}>
        <Input
          label="Prénom"
          name="firstName"
          id="firstName"
          type="text"
          classes="form__input"
          required={true}
          placeholder="Veuillez saisir votre prénom"
          handleChange={(e) => setUser({ ...user, firstName: e.target.value })}
        />
        <Input
          label="Nom"
          name="lastName"
          id="lastName"
          type="text"
          classes="form__input"
          required={true}
          placeholder="Veuillez saisir votre nom de famille"
          handleChange={(e) => setUser({ ...user, lastName: e.target.value })}
        />
        <Input
          label="Username"
          name="username"
          id="username"
          type="text"
          classes="form__input"
          required={true}
          placeholder="Veuillez saisir votre username"
          handleChange={(e) => setUser({ ...user, username: e.target.value })}
        />
        <Input
          label="Email"
          name="email"
          id="email"
          type="email"
          classes="form__input"
          required={true}
          placeholder="Veuillez saisir votre nom email"
          handleChange={(e) => setUser({ ...user, email: e.target.value })}
        />
        <Input
          label="Mot de passe"
          name="password"
          id="password"
          type="password"
          classes="form__input"
          required={true}
          placeholder="Veuillez saisir votre mot de passe"
          handleChange={(e) => setUser({ ...user, password: e.target.value })}
        />
        <Button title="envoyer" classes="btn btn__color-black" type="submit" />
      </form>
    </div>
  );
};

export default Index;
