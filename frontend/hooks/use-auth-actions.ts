"use client";
import { useQueryClient } from "@tanstack/react-query";
import { useRouter } from "next/navigation";

export type AuthType = "login" | "register";

type AuthPayload = {
    email: string;
    password: string;
    username?: string;
};

export const useAuthActions = () => {
    const router = useRouter();
    const queryClient = useQueryClient()

    const submitAuth = async (type: AuthType, payload: AuthPayload) => {
        const endpoint = type === "login" ? "token" : "register";

        try {
            const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/${endpoint}`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(payload),
                credentials: "include",
            });

            if (!response.ok) {
                const message = await response.text();
                throw new Error(message || "Erro ao autenticar");
            }

            router.push(type === "login" ? "/dashoard" : "/login");
        } catch (err) {
            throw err;
        }
    };

    const logout = async () => {
        await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/destroy-token`, {
            method: 'POST',
            credentials: 'include',
        });

        await queryClient.invalidateQueries({ queryKey: ['me'] });
        router.push('/login');
    };

    return { submitAuth, logout };
};

