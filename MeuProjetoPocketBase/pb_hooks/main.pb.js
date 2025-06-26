/// <reference path="../pb_data/types.d.ts" />

// Exemplo de hook customizado
onRecordAfterCreateRequest((e) => {
    console.log("Novo registro criado:", e.record);
    
    // Exemplo: enviar email quando um novo usuário se cadastra
    if (e.collection.name === "users") {
        // Lógica para enviar email de boas-vindas
        console.log("Enviando email de boas-vindas para:", e.record.email);
    }
}, "users");

// Exemplo de rota customizada
routerAdd("GET", "/api/custom/hello", (e) => {
    return e.json(200, {
        message: "Hello from custom route!",
        timestamp: new Date().toISOString()
    });
});

// Exemplo de validação customizada
onRecordBeforeCreateRequest((e) => {
    // Adicionar validações customizadas
    if (e.collection.name === "posts" && !e.record.title) {
        throw new Error("O título é obrigatório");
    }
}, "posts");

// Exemplo de middleware
onBeforeServe((e) => {
    console.log("PocketBase iniciado com customizações!");
}); 