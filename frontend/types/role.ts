export type RoleNames = 'admin' | 'user' | 'guest'

export type Role = {
    id: number
    name: RoleNames
    description: string
    level: number
}
