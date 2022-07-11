import { DateTime } from "luxon"

export interface TaskStatus {
    nextRunTime: DateTime
    id: number
    name: string
}
