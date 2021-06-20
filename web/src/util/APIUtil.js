import store from "../store";
import AuthUtil from "./AuthUtil";

const doAPIReq = (url, options, useAuth = true, errMsg = undefined, ) => {
  return new Promise((resolve, reject) => {
    if (useAuth) {
      if (!options.headers) {
        options.headers = {};
      }
      try {
        const tok = AuthUtil.getIDToken();
        options.headers["Authorization"] = tok;
      } catch (err) {
        console.log(err);
      }
    }
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
          if (errMsg) {
            store.commit("addToast", {
              type: "error",
              msg: errMsg
            })
          }
          reject(`api request error: `, resp.msg);
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
    false,
    "Authentication error"
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
    false,
    "Error registering user"
  );
};

const createNewProfile = (name) => {
  return doAPIReq(
    `/api/v1/profile`,
    {
      method: "POST",
      body: JSON.stringify({
        name,
      }),
    },
    true,
    `Error creating profile ${name}`
  );
};

const getProfiles = () => {
  return doAPIReq(
    `/api/v1/profile`,
    {
      method: "GET"
    },
    true,
    `Error getting profiles`
  )
}

const updateProfile = (id, name, filter, sorter) => {
  return doAPIReq(
    `/api/v1/profile/${id}`,
    {
      method: "PUT",
      body: JSON.stringify({
        id,
        name,
        filter,
        sorter
      })
    },
    true,
    `Error updating profile ${name}`
  )
}

const getConfiguration = () => {
  return doAPIReq(
    `/api/v1/configuration`,
    {
      method: "GET",
    },
    true,
  );
};

const getReleaseProfileCfg = () => {
  return doAPIReq(
    `/api/v1/releaseProfileCfg`,
    {
      method: "GET",
    },
    true,
  );
};

export default {
  signIn,
  signUp,
  isFirstTime,
  createNewProfile,
  getProfiles,
  updateProfile,
  getConfiguration,
  getReleaseProfileCfg,
};
