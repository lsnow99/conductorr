import AuthUtil from "./AuthUtil";

const doAPIReq = (url, options, useAuth = true) => {
  return new Promise((resolve, reject) => {
    fetch(url, options)
      .then((re) => re.json())
      .then((resp) => {
        if (resp.success) {
          if (resp.token) {
            AuthUtil.setIDToken(resp.token);
          }
          if (resp.data) {
            resolve(resp.data);
          } else {
            resolve({});
          }
        } else {
          console.log(resp.msg);
          reject(`api request error: ${resp.msg}`);
        }
      })
      .catch((err) => {
        reject(err);
      });
  });
};

const isFirstTime = () => {
  return doAPIReq(
    `/api/v1/first_time`,
    {
      method: "GET",
    },
    false
  );
};

const signIn = (username, password) => {
  return doAPIReq(
    `/api/v1/signin`,
    {
      method: "POST",
      body: JSON.stringify({
        username,
        password,
      }),
    },
    false
  );
};

const signUp = (username, password) => {
  return doAPIReq(
    `/api/v1/signup`,
    {
      method: "POST",
      body: JSON.stringify({
        username,
        password,
      }),
    },
    false
  );
};

export default {
  signIn,
  signUp,
  isFirstTime,
};
