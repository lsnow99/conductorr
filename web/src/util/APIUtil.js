import store from "../store";
import AuthUtil from "./AuthUtil";

const doAPIReq = (url, options, errMsg = undefined) => {
  return new Promise((resolve, reject) => {
    fetch(url, options)
      .then((re) => re.json())
      .then((resp) => {
        if (resp.success) {
          if (resp.data !== undefined) {
            resolve(resp.data);
          } else {
            resolve({});
          }
        } else {
          if (resp.failed_auth) {
            AuthUtil.logout();
            this.$oruga.notification.open({
              duration: 3000,
              message: `Authentication error`,
              position: "bottom-right",
              variant: "danger",
              closable: false,
            });
          } else if (errMsg) {
            this.$oruga.notification.open({
              duration: 3000,
              message: errMsg,
              position: "bottom-right",
              variant: "danger",
              closable: false,
            });
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
  return doAPIReq(`/api/v1/firstTime`, {
    method: "GET",
  });
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
    `Error creating profile ${name}`
  );
};

const searchLibrary = (query, contentType, page) => {
  return doAPIReq(
    `/api/v1/library/search?q=${query}&type=${contentType}&page=${page}`,
    {
      method: "GET",
    },
    `Error searching library`
  );
};

const searchNew = (query, contentType, page) => {
  return doAPIReq(
    `/api/v1/new/search?q=${query}&type=${contentType}&page=${page}`,
    {
      method: "GET",
    },
    `Error searching for media`
  );
};

const getProfiles = () => {
  return doAPIReq(
    `/api/v1/profile`,
    {
      method: "GET",
    },
    `Error getting profiles`
  );
};

const updateProfile = (id, name, filter, sorter) => {
  return doAPIReq(
    `/api/v1/profile/${id}`,
    {
      method: "PUT",
      body: JSON.stringify({
        id,
        name,
        filter,
        sorter,
      }),
    },
    `Error updating profile ${name}`
  );
};

const deleteProfile = (id) => {
  return doAPIReq(
    `/api/v1/profile/${id}`,
    {
      method: "DELETE",
    },
    `Error deleting profile`
  );
};

const getConfiguration = () => {
  return doAPIReq(`/api/v1/configuration`, {
    method: "GET",
  });
};

const getReleaseProfileCfg = () => {
  return doAPIReq(`/api/v1/releaseProfileCfg`, {
    method: "GET",
  });
};

const addMedia = (imdb_id) => {
  return doAPIReq(
    `/api/v1/add/${imdb_id}`,
    {
      method: "POST",
    },
    `Error adding media`
  );
};

const checkAuth = () => {
  return doAPIReq(`/api/v1/checkAuth`, {
    method: "GET",
  });
};

const getMedia = (id) => {
  return doAPIReq(`/api/v1/media/${id}`, {
    method: "GET"
  })
}

const getProfile = (id) => {
  return doAPIReq(`/api/v1/profile/${id}`, {
    method: "GET"
  })
}

export default {
  signIn,
  signUp,
  isFirstTime,
  createNewProfile,
  getProfiles,
  updateProfile,
  deleteProfile,
  getConfiguration,
  getReleaseProfileCfg,
  searchLibrary,
  searchNew,
  addMedia,
  checkAuth,
  getMedia,
  getProfile
};
