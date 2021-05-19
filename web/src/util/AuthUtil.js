import jwt_decode from "jwt-decode";
import store from "../store";

const logout = () => {
  localStorage.clear();
};

const getIDToken = () => {
    const idTok = localStorage.getItem("id_token")
    return new Promise((resolve, reject) => {
        if (idTok) {
            resolve(idTok)
        } else {
            reject("not logged in")
        }
    })
}

const setIDToken = (idTok) => {
    localStorage.setItem("id_token", idTok)
}

const getLoggedInID = () => {
  return new Promise((resolve, reject) => {
    getIDToken()
      .then((jwt) => {
        store.commit('setLoggedIn', true)
        resolve(jwt_decode(jwt).sub);
      })
      .catch((err) => {
        store.commit('setLoggedIn', false)
        reject(err);
      });
  });
};

export default {
  logout,
  getLoggedInID,
  setIDToken
};
