import jwt_decode from "jwt-decode";
import store from "../store";

const logout = () => {
  localStorage.clear();
};

const getIDToken = () => {
    const idTok = localStorage.getItem("id_token")
    if (!idTok) {
      throw "not logged in"
    }
    return idTok
}

const setIDToken = (idTok) => {
  localStorage.setItem("id_token", idTok)
  store.commit('setLoggedIn', true)
}

const getLoggedInID = () => {
  return new Promise((resolve, reject) => {
    try {
      let idTok = getIDToken()
      store.commit('setLoggedIn', true)
      resolve(jwt_decode(idTok).sub)
    } catch (err) {
      store.commit('setLoggedIn', false)
      console.log(err)
      reject(err)
    }
  });
};

export default {
  logout,
  getLoggedInID,
  getIDToken,
  setIDToken
};
