export interface StatusMessage {
    system_name: string
    msg: string
}

export interface Status {
    statuses: StatusMessage[]
    healthy: boolean
    msg: string
}