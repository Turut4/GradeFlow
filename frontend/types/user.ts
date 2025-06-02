import { Role } from "./role"

export type User = {
    id: number
    username: string
    email: string
    password: string
    role: Role
    roleId: number
    deletedAt?: string | null
    updatedAt: string
    createdAt: string
} 
