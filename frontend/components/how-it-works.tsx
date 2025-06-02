import { Card, CardContent } from "@/components/ui/card"
import { Upload, Scan, CheckCircle, BarChart } from "lucide-react"

const steps = [
    {
        icon: Upload,
        title: "Faça Upload",
        description: "Tire uma foto do gabarito ou faça upload da imagem diretamente na plataforma.",
        step: "01",
    },
    {
        icon: Scan,
        title: "Processamento IA",
        description: "Nossa inteligência artificial analisa e reconhece automaticamente as marcações.",
        step: "02",
    },
    {
        icon: CheckCircle,
        title: "Correção Automática",
        description: "O sistema compara com o gabarito oficial e gera a correção instantaneamente.",
        step: "03",
    },
    {
        icon: BarChart,
        title: "Relatórios",
        description: "Receba relatórios detalhados com análises de desempenho e estatísticas.",
        step: "04",
    },
]

export default function HowItWorks() {
    return (
        <section id="sobre" className="py-20 bg-gray-50">
            <div className="container mx-auto px-4">
                <div className="text-center mb-16">
                    <h2 className="text-3xl font-bold text-gray-900 sm:text-4xl mb-4">Como Funciona</h2>
                    <p className="text-lg text-gray-600 max-w-2xl mx-auto">
                        Em apenas 4 passos simples, transforme a correção manual em um processo automatizado e eficiente.
                    </p>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
                    {steps.map((step, index) => (
                        <Card key={index} className="relative border-0 shadow-lg bg-white">
                            <CardContent className="p-8 text-center">
                                <div className="absolute -top-4 left-1/2 transform -translate-x-1/2">
                                    <div className="w-8 h-8 bg-blue-600 text-white rounded-full flex items-center justify-center text-sm font-bold">
                                        {step.step}
                                    </div>
                                </div>
                                <div className="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-6 mt-4">
                                    <step.icon className="w-8 h-8 text-blue-600" />
                                </div>
                                <h3 className="text-xl font-semibold text-gray-900 mb-3">{step.title}</h3>
                                <p className="text-gray-600 leading-relaxed">{step.description}</p>
                            </CardContent>
                        </Card>
                    ))}
                </div>
            </div>
        </section>
    )
}

