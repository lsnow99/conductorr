import { doAPIReq } from "./API"

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

const addMedia = (searchID, profileID, pathID) => {
  return doAPIReq(
    `/api/v1/add/${searchID}`,
    {
      method: "POST",
      body: JSON.stringify({
        profileID,
        pathID
      })
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

const searchReleasesAuto = (id) => {
  return doAPIReq(`/api/v1/media/${id}/searchReleasesAuto`, {
    method: "POST"
  })
}

const testIndexer = (name, baseUrl, apiKey, forMovies, forSeries, downloadType) => {
  return doAPIReq(`/api/v1/testIndexer`, {
    method: "POST",
    body: JSON.stringify({
      name,
      baseUrl,
      apiKey,
      forMovies,
      forSeries,
      downloadType
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

const updateMedia = (id, profile_id, path_id) => {
  return doAPIReq(`/api/v1/media/${id}`, {
    method: "PUT",
    body: JSON.stringify({
      profile_id,
      path_id
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

const newDownloader = (downloader_type, name, file_action, config) => {
  return doAPIReq(`/api/v1/downloader`, {
    method: "POST",
    body: JSON.stringify({
      name,
      downloader_type,
      file_action,
      config
    })
  })
}

const getDownloaders = () => {
  return doAPIReq(`/api/v1/downloader`, {
    method: "GET",
  })
}

const updateDownloader = (id, name, file_action, config) => {
  console.log('id', id, 'name', name, 'file_action', file_action, 'config', config)
  return doAPIReq(`/api/v1/downloader/${id}`, {
    method: "PUT",
    body: JSON.stringify({
      name,
      file_action,
      config
    })
  })
}

const deleteDownloader = (id) => {
  return doAPIReq(`/api/v1/downloader/${id}`, {
    method: "DELETE",
  })
}

const downloadMediaRelease = (media_id, release) => {
  return doAPIReq(`/api/v1/media/${media_id}/download`, {
    method: "POST",
    body: JSON.stringify(release)
  })
}

const getStatus = () => {
  return doAPIReq(`/api/v1/status`, {
    method: "GET",
  })
}

const getLogs = () => {
  return doAPIReq(`/api/v1/logs`, {
    method: "GET"
  })
}

const createNewPath = (path, movies_default, series_default) => {
  return doAPIReq(`/api/v1/path`, {
    method: "POST",
    body: JSON.stringify({
      path,
      movies_default,
      series_default
    })
  })
}

const updatePath = (id, path, movies_default, series_default) => {
  return doAPIReq(`/api/v1/path/${id}`, {
    method: "PUT",
    body: JSON.stringify({
      path,
      movies_default,
      series_default
    })
  })
}

const deleteMedia = (id) => {
  return doAPIReq(`/api/v1/media/${id}`, {
    method: "DELETE",
  })
}

const getSchedule = () => {
  return doAPIReq(`/api/v1/schedule`, {
    method: "GET",
  })
}

const logout = () => {
  return doAPIReq(`/api/v1/logout`, {
    method: "GET"
  })
}

const setMonitoringMedia = (id, monitoring) => {
  return doAPIReq(`/api/v1/media/${id}/monitoring`, {
    method: "PUT",
    body: JSON.stringify({
      monitoring
    })
  })
}

const getActiveDownloads = (mediaID = 0) => {
  return doAPIReq(`/api/v1/activeDownloads?media_id=${mediaID}`, {
    method: "GET",
  })
}

const getFinishedDownloads = (mediaID = 0) => {
  return doAPIReq(`/api/v1/finishedDownloads?media_id=${mediaID}`, {
    method: "GET",
  })
}

const getPlexAuthToken = (username, password) => {
  return doAPIReq(`/api/v1/plexAuthToken`, {
    method: "POST",
    body: JSON.stringify({
      username,
      password
    })
  })
}

const testMediaServer = (media_server_type, config) => {
  return doAPIReq(`/api/v1/testMediaServer`, {
    method: "POST",
    body: JSON.stringify({
      media_server_type,
      config,
    })
  })
}

const newMediaServer = (media_server_type, name, config) => {
  return doAPIReq(`/api/v1/mediaServer`, {
    method: "POST",
    body: JSON.stringify({
      name,
      media_server_type,
      config
    })
  })
}

const getMediaServers = () => {
  return doAPIReq(`/api/v1/mediaServer`, {
    method: "GET",
  })
}

const updateMediaServer = (id, name, config) => {
  return doAPIReq(`/api/v1/mediaServer/${id}`, {
    method: "PUT",
    body: JSON.stringify({
      name,
      config
    })
  })
}

const deleteMediaServer = (id) => {
  return doAPIReq(`/api/v1/mediaServer/${id}`, {
    method: "DELETE",
  })
}

const refreshMediaMetadata = (id) => {
  return doAPIReq(`/api/v1/media/${id}/refresh`, {
    method: "POST",
  })
}

const createNewBackup = () => {
  return doAPIReq(`/api/v1/backup`, {
    method: "POST"
  })
}

const getTaskStatuses = () => {
  return doAPIReq(`/api/v1/tasks`, {
    method: "GET"
  })
}

const getRecentMedia = () => {
  return doAPIReq(`/api/v1/library/recent`, {
    method: "GET"
  })
}

const downloadLogs = () => {
  return doAPIReq(`/api/v1/logs`, {
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
  getProfile,
  searchReleasesManual,
  searchReleasesAuto,
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
  getStatus,
  getLogs,
  createNewPath,
  updatePath,
  deleteMedia,
  getSchedule,
  logout,
  setMonitoringMedia,
  getActiveDownloads,
  getFinishedDownloads,
  getPlexAuthToken,
  testMediaServer,
  newMediaServer,
  getMediaServers,
  updateMediaServer,
  deleteMediaServer,
  refreshMediaMetadata,
  createNewBackup,
  getTaskStatuses,
  getRecentMedia,
  downloadLogs
};
