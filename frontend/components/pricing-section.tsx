import { Button } from '@/components/ui/button';
import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
} from '@/components/ui/card';
import { useAuth } from '@/context/auth-context';
import { Check, Star } from 'lucide-react';
import Link from 'next/link';

const plans = [
	{
		name: 'Gratuito',
		price: 'R$ 0',
		period: '/mês',
		description: 'Perfeito para começar',
		features: [
			'Até 50 correções por mês',
			'Suporte básico por email',
			'Relatórios simples',
			'Armazenamento por 30 dias',
		],
		cta: 'Começar Grátis',
		popular: false,
	},
	{
		name: 'Educador',
		price: 'R$ 29',
		period: '/mês',
		description: 'Para professores ativos',
		features: [
			'Até 500 correções por mês',
			'Suporte prioritário',
			'Relatórios avançados',
			'Gestão de turmas',
			'Armazenamento ilimitado',
			'Exportação de dados',
		],
		cta: 'Assinar Agora',
		popular: true,
	},
	{
		name: 'Instituição',
		price: 'R$ 99',
		period: '/mês',
		description: 'Para escolas e cursos',
		features: [
			'Correções ilimitadas',
			'Suporte 24/7',
			'Relatórios personalizados',
			'Múltiplos usuários',
			'API de integração',
			'Treinamento personalizado',
			'Backup automático',
		],
		cta: 'Falar com Vendas',
		popular: false,
	},
];

export default function PricingSection() {
	return (
		<section id="precos" className="py-20 bg-white">
			<div className="container mx-auto px-4">
				<div className="text-center mb-16">
					<h2 className="text-3xl font-bold text-gray-900 sm:text-4xl mb-4">
						Planos para Todos os Tamanhos
					</h2>
					<p className="text-lg text-gray-600 max-w-2xl mx-auto">
						Escolha o plano ideal para suas necessidades. Comece grátis e evolua
						conforme sua demanda cresce.
					</p>
				</div>

				<div className="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-6xl mx-auto">
					{plans.map((plan, index) => (
						<Card
							key={index}
							className={`relative border-2 ${
								plan.popular
									? 'border-blue-500 shadow-xl scale-105'
									: 'border-gray-200 shadow-lg'
							}`}
						>
							{plan.popular && (
								<div className="absolute -top-4 left-1/2 transform -translate-x-1/2">
									<div className="bg-blue-500 text-white px-4 py-1 rounded-full text-sm font-medium flex items-center">
										<Star className="w-4 h-4 mr-1" />
										Mais Popular
									</div>
								</div>
							)}

							<CardHeader className="text-center pb-8 pt-8">
								<CardTitle className="text-2xl font-bold text-gray-900">
									{plan.name}
								</CardTitle>
								<CardDescription className="text-gray-600 mt-2">
									{plan.description}
								</CardDescription>
								<div className="mt-6">
									<span className="text-4xl font-bold text-gray-900">
										{plan.price}
									</span>
									<span className="text-gray-600">{plan.period}</span>
								</div>
							</CardHeader>

							<CardContent className="space-y-4">
								<ul className="space-y-3">
									{plan.features.map((feature, featureIndex) => (
										<li key={featureIndex} className="flex items-center">
											<Check className="w-5 h-5 text-green-500 mr-3 flex-shrink-0" />
											<span className="text-gray-600">{feature}</span>
										</li>
									))}
								</ul>

								<div className="pt-6">
									<Button
										className={`w-full ${
											plan.popular
												? 'bg-blue-600 hover:bg-blue-700'
												: 'bg-gray-900 hover:bg-gray-800'
										}`}
										asChild
									>
										<Link href="/cadastro">{plan.cta}</Link>
									</Button>
								</div>
							</CardContent>
						</Card>
					))}
				</div>

				<div className="text-center mt-12">
					<p className="text-gray-600">
						Todos os planos incluem teste grátis de 14 dias. Cancele a qualquer
						momento.
					</p>
				</div>
			</div>
		</section>
	);
}
