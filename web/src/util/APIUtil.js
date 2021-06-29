import store from "../store";
import AuthUtil from "./AuthUtil";
import EventBus from "./EventBus";

const doAPIReq = (url, options, errMsg = undefined) => {
  return new Promise((resolve, reject) => {
    fetch(url, options)
      .then((re) => re.json())
      .then((resp) => {
        console.log(resp)
        if (resp.success) {
          if (resp.data !== undefined) {
            resolve(resp.data);
          } else {
            resolve({});
          }
        } else {
          if (resp.failed_auth) {
            AuthUtil.logout();
            EventBus.emit('notification', {
              duration: 3000,
              message: `Authentication error`,
              position: "bottom-right",
              variant: "danger",
              closable: false,
            })
          } else if (errMsg) {
            EventBus.emit('notification', {
              duration: 3000,
              message: errMsg,
              position: "bottom-right",
              variant: "danger",
              closable: false,
            })
          }
          reject(resp.msg);
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

const searchReleasesManual = (id) => {
  return doAPIReq(`/api/v1/media/${id}/searchReleasesManual`, {
    method: "GET"
  })
}

const testIndexer = (name, base_url, api_key, for_movies, for_series, download_type) => {
  return doAPIReq(`/api/v1/testIndexer`, {
    method: "POST",
    body: JSON.stringify({
      name,
      base_url,
      api_key,
      for_movies,
      for_series,
      download_type
    })
  })
}

const newIndexer = (name, base_url, api_key, for_movies, for_series, download_type) => {
  return doAPIReq(`/api/v1/indexer`, {
    method: "POST",
    body: JSON.stringify({
      name,
      base_url,
      api_key,
      for_movies,
      for_series,
      download_type
    })
  })
}

const getIndexers = () => {
  return doAPIReq(`/api/v1/indexer`, {
    method: "GET",
  })
}

const updateIndexer = (id, name, base_url, api_key, for_movies, for_series, download_type) => {
  return doAPIReq(`/api/v1/indexer/${id}`, {
    method: "PUT",
    body: JSON.stringify({
      name,
      base_url,
      api_key,
      for_movies,
      for_series,
      download_type
    })
  })
}

const deleteIndexer = (id) => {
  return doAPIReq(`/api/v1/indexer/${id}`, {
    method: "DELETE"
  })
}

const testPath = (path) => {
  return doAPIReq(`/api/v1/testPath`, {
    method: "POST",
    body: JSON.stringify({
      path
    })
  })
}

const updatePaths = (paths) => {
  return doAPIReq(`/api/v1/path`, {
    method: "PUT",
    body: JSON.stringify(paths)
  })
}

const getPaths = () => {
  return doAPIReq(`/api/v1/path`, {
    method: "GET",
  })
}

const deletePath = (id) => {
  return doAPIReq(`/api/v1/path/${id}`, {
    method: "DELETE"
  })
}

const updateMedia = (id, profile_id) => {
  return doAPIReq(`/api/v1/media/${id}`, {
    method: "PUT",
    body: JSON.stringify({
      profile_id
    })
  })
}

const testDownloader = (downloader_type, config) => {
  return doAPIReq(`/api/v1/testDownloader`, {
    method: "POST",
    body: JSON.stringify({
      downloader_type,
      config,
    })
  })
}

const newDownloader = (downloader_type, name, config) => {
  return doAPIReq(`/api/v1/downloader`, {
    method: "POST",
    body: JSON.stringify({
      name,
      downloader_type,
      config
    })
  })
}

const getDownloaders = () => {
  return doAPIReq(`/api/v1/downloader`, {
    method: "GET",
  })
}

const updateDownloader = (id, name, config) => {
  return doAPIReq(`/api/v1/downloader/${id}`, {
    method: "PUT",
    body: JSON.stringify({
      name,
      config
    })
  })
}

const deleteDownloader = (id) => {
  return doAPIReq(`/api/v1/downloader/${id}`, {
    method: "DELETE",
  })
}

const downloadMediaRelease = (mediaID, download_url) => {
  return doAPIReq(`/api/v1/media/${mediaID}/download`, {
    method: "POST",
    body: JSON.stringify({
      download_url
    })
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
  getProfile,
  searchReleasesManual,
  testIndexer,
  newIndexer,
  getIndexers,
  updateIndexer,
  testPath,
  deleteIndexer,
  updatePaths,
  getPaths,
  deletePath,
  updateMedia,
  testDownloader,
  newDownloader,
  getDownloaders,
  updateDownloader,
  deleteDownloader,
  downloadMediaRelease,
};
