import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Zap, BarChart3, Shield, Users, Camera, FileText } from "lucide-react"

const features = [
    {
        icon: Camera,
        title: "Escaneamento Inteligente",
        description:
            "Capture gabaritos com a câmera do celular ou faça upload de imagens. Nossa IA reconhece automaticamente as marcações.",
    },
    {
        icon: Zap,
        title: "Correção Instantânea",
        description:
            "Resultados em segundos com precisão de 99.9%. Identifica automaticamente respostas corretas e incorretas.",
    },
    {
        icon: BarChart3,
        title: "Relatórios Detalhados",
        description: "Análises completas de desempenho, estatísticas por questão e identificação de pontos de melhoria.",
    },
    {
        icon: Users,
        title: "Gestão de Turmas",
        description: "Organize alunos em turmas, acompanhe o progresso individual e coletivo de forma simples.",
    },
    {
        icon: FileText,
        title: "Múltiplos Formatos",
        description: "Suporte para diferentes tipos de gabarito: múltipla escolha, verdadeiro/falso e questões numéricas.",
    },
    {
        icon: Shield,
        title: "Segurança Total",
        description: "Dados criptografados e armazenamento seguro. Conformidade com LGPD e proteção da privacidade.",
    },
]

export default function FeaturesSection() {
    return (
        <section id="recusros" className="py-20 bg-white">
            <div className="container mx-auto px-4">
                <div className="text-center mb-16">
                    <h2 className="text-3xl font-bold text-gray-900 sm:text-4xl mb-4">Recursos Poderosos para Educadores</h2>
                    <p className="text-lg text-gray-600 max-w-2xl mx-auto">
                        Tudo que você precisa para modernizar o processo de correção e focar no que realmente importa: o aprendizado
                        dos seus alunos.
                    </p>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                    {features.map((feature, index) => (
                        <Card key={index} className="border-0 shadow-lg hover:shadow-xl transition-shadow duration-300">
                            <CardHeader>
                                <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mb-4">
                                    <feature.icon className="w-6 h-6 text-blue-600" />
                                </div>
                                <CardTitle className="text-xl font-semibold text-gray-900">{feature.title}</CardTitle>
                            </CardHeader>
                            <CardContent>
                                <CardDescription className="text-gray-600 text-base leading-relaxed">
                                    {feature.description}
                                </CardDescription>
                            </CardContent>
                        </Card>
                    ))}
                </div>
            </div>
        </section>
    )
}

