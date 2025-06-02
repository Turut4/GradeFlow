'use client';

import { useState, useCallback } from 'react';
import { Button } from '@/components/ui/button';
import { AuthType, useAuthActions } from '@/hooks/use-auth-actions';

export default function AuthForm({ type }: { type: AuthType }) {
    const [formData, setFormData] = useState({
        username: '',
        email: '',
        password: '',
        confirmPassword: '',
    });
    const [isSubmitting, setIsSubmitting] = useState(false);
    const { submitAuth } = useAuthActions();

    const handleInputChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData((prev) => ({ ...prev, [name]: value }));
    }, []);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        if (type === 'register' && formData.password !== formData.confirmPassword) {
            alert('As senhas não coincidem');
            setIsSubmitting(false);
            return;
        }

        try {
            await submitAuth(type, {
                email: formData.email,
                password: formData.password,
                ...(type === 'register' && { username: formData.username }),
            });
        } catch (error) {
            alert(
                `Falha ao ${type === 'login' ? 'entrar' : 'registrar'}: ${error instanceof Error ? error.message : 'Erro desconhecido'
                }`
            );
        } finally {
            setIsSubmitting(false);
        }
    };


    return (
        <form
            onSubmit={handleSubmit}
            className="space-y-6 bg-white -border-2 p-15 rounded-2xl shadow-white"
            aria-labelledby="form-heading"
        >
            <h1 id="form-heading" className="text-2xl font-bold text-center">
                {type === 'login' ? 'Entrar na Plataforma' : 'Criar Nova Conta'}
            </h1>

            <div className="space-y-4">
                {type === 'register' && (
                    <InputField
                        id="username"
                        label="Nome de usuário"
                        name="username"
                        value={formData.username}
                        onChange={handleInputChange}
                        placeholder="Seu nome de usuário"
                        autoComplete="username"
                        required
                    />
                )}

                <InputField
                    id="email"
                    type="email"
                    label="E-mail"
                    name="email"
                    value={formData.email}
                    onChange={handleInputChange}
                    placeholder="seu@email.com"
                    autoComplete="email"
                    required
                />

                <InputField
                    id="password"
                    type="password"
                    label="Senha"
                    name="password"
                    value={formData.password}
                    onChange={handleInputChange}
                    placeholder="••••••••"
                    autoComplete={type === 'login' ? 'current-password' : 'new-password'}
                    required
                />

                {type === 'register' && (
                    <InputField
                        id="confirmPassword"
                        type="password"
                        label="Confirmar Senha"
                        name="confirmPassword"
                        value={formData.confirmPassword}
                        onChange={handleInputChange}
                        placeholder="••••••••"
                        autoComplete="new-password"
                        required
                    />
                )}
            </div>

            <Button
                type="submit"
                className="w-full"
                disabled={isSubmitting}
                aria-busy={isSubmitting}
            >
                {isSubmitting
                    ? 'Processando...'
                    : type === 'login'
                        ? 'Entrar'
                        : 'Registrar'}
            </Button>
        </form>
    );
}

const InputField = ({
    id,
    label,
    type = 'text',
    name,
    value,
    onChange,
    placeholder,
    autoComplete,
    required,
}: {
    id: string;
    label: string;
    type?: string;
    name: string;
    value: string;
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
    placeholder: string;
    autoComplete?: string;
    required?: boolean;
}) => (
    <div>
        <label htmlFor={id} className="block text-sm font-medium mb-1">
            {label}
            {required && <span className="text-red-500">*</span>}
        </label>
        <input
            id={id}
            type={type}
            name={name}
            value={value}
            onChange={onChange}
            placeholder={placeholder}
            autoComplete={autoComplete}
            required={required}
            className="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
            aria-required={required}
        />
    </div>
);
