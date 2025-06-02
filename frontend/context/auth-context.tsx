'use client';

import { createContext, useContext, useMemo } from 'react';
import { useQuery } from '@tanstack/react-query';
import { User } from '@/types/user';

async function fetchUser(): Promise<User | null> {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/users/me`, {
        credentials: 'include',
    });

    if (res.ok) {
        const json: { data: User } = await res.json();
        return json.data;
    }

    return null;
}

interface AuthContextType {
    user: User | null | undefined;
    isLoading: boolean;
    isAuthenticated: boolean;
    refetch: ReturnType<typeof useQuery>['refetch'];
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
    const {
        data: user,
        isLoading,
        refetch,
    } = useQuery({
        queryKey: ['me'],
        queryFn: fetchUser,
        staleTime: 1000 * 60 * 5,
        retry: false,
    });

    const isAuthenticated = !!user;
    const value = useMemo(
        () => ({
            user,
            isLoading,
            refetch,
            isAuthenticated,
        }),
        [user, isLoading, isAuthenticated, refetch]
    );
    return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error('useAuth deve ser usado dentro de um <AuthProvider>');
    }
    return context;
};
