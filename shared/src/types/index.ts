import { DateTime } from "luxon";

export enum Variant {
    SUCCESS = "success",
    WARNING = "warning",
    DANGER = "danger",
    INFO = "info"
}

export interface LogMessage {
    variant: Variant
    msg: string
    timestamp: DateTime
    decoration?: string
}