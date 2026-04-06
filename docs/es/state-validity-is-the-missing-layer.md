# MCP no es suficiente: la capa que falta es State Validity

La infraestructura de IA está convergiendo rápidamente hacia un nuevo consenso.

Más protocolos.  
Más herramientas.  
Más evaluaciones.  
Más interoperabilidad.  
Más flujos de trabajo con agentes.  

Todo eso importa.  
Y todo eso tiene valor.

Pero creo que a la industria todavía le falta una capa más importante:

**state validity**

Que un sistema esté bien conectado no significa que tenga legitimidad para actuar.

Puede tener el modelo correcto,  
el acceso correcto a herramientas,  
las instrucciones correctas del repositorio,  
el benchmark correcto,  
la integración correcta del workflow,  

y aun así estar operando desde un **estado incorrecto**.

Puede estar leyendo un handoff obsoleto.  
Puede estar actuando desde un marco de autoridad ya vencido.  
Puede estar arrastrando permisos que ya deberían haber expirado.  
Puede estar tratando trazas históricas como si fueran autorización actual.  
Puede estar confundiendo una superficie de entrada vieja con la que realmente está viva.  
O puede estar tomando una sincronización parcial como si fuera un presente válido.

Eso no es solo un problema de capability.

Es un problema de **state validity**.

Y state validity se convertirá en uno de los conceptos más importantes para los sistemas de IA serios.

Porque la verdadera pregunta ya no es solo:

**¿Puede el sistema hacer la tarea?**

La pregunta más difícil es:

**¿Este estado exacto del sistema todavía tiene el derecho de hablar, decidir o modificar el mundo formal ahora mismo?**

Ese listón es mucho más alto.

Y debería serlo.

Estamos entrando en una etapa en la que  
los modelos pueden llamar herramientas,  
las herramientas pueden activar workflows,  
los workflows pueden tocar sistemas de producción,  
y múltiples agentes pueden coordinarse dentro de bucles.

En ese mundo, los mayores fallos no vendrán solo de un mal razonamiento.

Vendrán de la **continuación inválida**.

Un estado vencido que sigue actuando.  
Una salida de helper que fue sobrepromovida.  
Un handoff packet que parece actual pero no lo es.  
Una tarea que debía quedarse en candidate-only pero se deslizó hacia execution.  
Una traza que debía permanecer histórica pero vuelve a usarse como permiso vivo.

Por eso creo que la siguiente capa real después de la conectividad no es “más autonomía”.

Es **admissibility**.

Antes de que un sistema de IA actúe, debe existir una forma de juzgar si el estado actual sigue siendo admissible for action.

No si el modelo es inteligente.  
No si la herramienta está disponible.  
No si el benchmark se ve fuerte.  
No si el protocolo es elegante.

Sino si **el estado sigue siendo válido**.

Eso implica al menos cinco cosas.

Primero, **landing no es permission**.  
Escribir algo en un formal host no significa automáticamente que ya esté autorizado para cambiar el mundo.

Segundo, **adoption no es execution**.  
Un sistema puede reconocer, absorber o registrar un objeto sin estar autorizado a actuar sobre él.

Tercero, **execution no es clearance permanente**.  
Incluso una acción válida no crea legitimidad infinita para seguir actuando.

Cuarto, **history no es current authority**.  
Las trazas, logs y packets anteriores pueden explicar el pasado, pero no autorizan automáticamente el presente.

Quinto, **el estado sincronizado más reciente debe prevalecer sobre superficies de entrada más antiguas**.  
Si varias capas de entrada entran en conflicto, el sistema debe decidir primero qué es realmente current antes de permitir que el opener defina el siguiente movimiento.

Aquí es donde gran parte de la infraestructura de agentes todavía se siente incompleta.

Estamos mejorando en conexión.  
Mejorando en orquestación.  
Mejorando en evaluación.  
Mejorando en tool calling.  

Pero todavía no tenemos suficientes sistemas que pregunten, antes de actuar:

**¿Este presente sigue siendo válido?**

Eso suena abstracto.

No lo es en absoluto.

Decide si un coding agent puede tocar un repo.  
Si un workflow puede continuar después de un checkpoint.  
Si un handoff puede definir la siguiente tarea.  
Si un helper agent debe permanecer candidate-only.  
Si una execution trace debe demobilize antes de que pueda empezar algo nuevo.  
Si un automated loop sigue siendo un lawful probe o ya derivó hacia unauthorized action.

En otras palabras:

**state validity es el punto donde governance deja de ser un documento de políticas y se convierte en un runtime substrate.**

Por eso creo que la infraestructura de IA de la próxima década se dividirá en dos grupos.

Un grupo seguirá optimizando throughput:  
más herramientas, más llamadas, más bucles, más cadenas autónomas.

El otro grupo construirá una capacidad más difícil:  
determinar con precisión bajo qué condiciones un sistema todavía conserva el derecho a actuar.

El primer grupo se moverá más rápido.  
El segundo grupo ganará más confianza.

Y trust, al final, se compone mejor que la velocidad bruta.

Así que sí:  
los protocolos importan.  
Las evaluaciones importan.  
Las superficies de instrucción importan.  
La interoperabilidad importa.

Pero ninguna de ellas es la última capa.

Porque un agente conectado no se convierte automáticamente en un actor válido.

La capa que falta es **state validity**.

Y creo que esa capa se convertirá en una de las abstracciones fundamentales de la próxima generación de sistemas de IA.

**MCP le dio a la IA un socket.  
State validity decidirá si tiene derecho a tocar el mundo.**
