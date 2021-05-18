import AuthUtil from "./AuthUtil";

const doAPIReq = (url, options, useAuth = true) => {
  return new Promise((resolve, reject) => {
    fetch(url, options)
      .then((re) => {
        if (re.status != 200) {
          reject(`unexpected status code: ${re.status}`);
        } else {
          const resp = JSON.parse(re.body);
          if (resp.success) {
            resolve(resp.data);
          } else {
            reject(`api request error: ${resp.msg}`);
          }
        }
      })
      .catch((err) => {
        reject(err);
      });
  });
};

const signIn = (username, password) => {
  return new Promise((resolve, reject) => {
    doAPIReq(`/api/v1/signin`, {
      method: "POST",
      body: JSON.stringify({
        username,
        password,
      }),
    })
      .then((data) => {
        AuthUtil.setIDToken(data.id_token);
        AuthUtil.setRefreshToken(data.refresh_token);
        resolve(data);
      })
      .catch((err) => {
        reject(err);
      });
  });
};

export default {
  signIn,
};
