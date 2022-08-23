import React from "react";

const Modal = (props) => {
  return (
    <>
      {props.isActive ? (
        <>
          <div className="overlay" onClick={props.closeFunction}></div>
          <div className="modal">
            <div className="modal__header">
              <div className="close__cross" onClick={props.closeFunction}>
                <span></span>
                <span></span>
              </div>
            </div>
            <div className="modal__title">
              <h2>{props.title}</h2>
            </div>
            <div className="modal__content">{props.children}</div>
            {props.type !== "information" ? (
              <div className="modal__actions">
                <button
                  className="btn btn__color-white"
                  onClick={props.closeFunction}
                >
                  Annuler
                </button>
                <button
                  className="btn btn__color-black"
                  onClick={props.validateFunction}
                >
                  Valider
                </button>
              </div>
            ) : (
              ""
            )}
          </div>
        </>
      ) : (
        ""
      )}
    </>
  );
};

export default Modal;
