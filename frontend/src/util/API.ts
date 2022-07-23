import AuthUtil from "./AuthUtil";
import EventBus from "./EventBus";

export const doAPIReq = async (
  url: string,
  options?: RequestInit | undefined,
  errMsg = undefined
) => {
  try {
    const re = await (await fetch(url, options)).json();
    if (re.success) {
      return re.data ?? {};
    } else {
      if (re.failed_auth) {
        AuthUtil.logout();
        EventBus.emit("forceLogout");
      } else if (errMsg) {
        EventBus.emit("notification", {
          duration: 3000,
          message: errMsg,
          position: "bottom-right",
          variant: "danger",
          closable: false,
        });
      }
      throw re.msg;
    }
  } catch (err) {
    throw err;
  }
};

export const getSchedule = (dateFrom: number, dateTo: number) =>
  doAPIReq(`/api/v1/schedule?dateFrom=${Math.round(dateFrom)}&dateTo=${Math.round(dateTo)}`, {
    method: "GET",
  });
