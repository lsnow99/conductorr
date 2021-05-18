import jwt_decode from "jwt-decode";

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

const setRefreshToken = (refreshTok) => {
    localStorage.setItem("refresh_token", refreshTok)
}

const getLoggedInID = () => {
  return new Promise((resolve, reject) => {
    getIDToken()
      .then((jwt) => {
        resolve(jwt_decode(jwt).sub);
      })
      .catch((err) => {
        reject(err);
      });
  });
};

export default {
  logout,
  getLoggedInID,
  setIDToken,
  setRefreshToken
};
