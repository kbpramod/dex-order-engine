D:.
│   .dockerignore
│   .gitignore
│   docker-compose.yml
│   Dockerfile
│   jest.config.js
│   package-lock.json
│   package.json
│   prisma.config.ts
│   README.md
│   tsconfig.json
│   
├───assets
│       architecture.png
│       jest.png
│       k6-limit.png
│       k6-market.png
│       k6-sniper.png
│       k6-ws.png
│       
├───generated
│   └───prisma
│       │   client.d.ts
│       │   client.js
│       │   default.d.ts
│       │   default.js
│       │   edge.d.ts
│       │   edge.js
│       │   index-browser.js
│       │   index.d.ts
│       │   index.js
│       │   package.json
│       │   query_compiler_bg.js
│       │   query_compiler_bg.wasm
│       │   query_compiler_bg.wasm-base64.js
│       │   schema.prisma
│       │   wasm-edge-light-loader.mjs
│       │   wasm-worker-loader.mjs
│       │   
│       └───runtime
│               client.d.ts
│               client.js
│               index-browser.d.ts
│               index-browser.js
│               wasm-compiler-edge.js
│               
├───k6
│   │   docker-compose.yml
│   │   grafana.json
│   │   limit_load_test.js
│   │   market_load_test.js
│   │   sniper_load_test.js
│   │   websocket_test.js
│   │   
│   └───prometheus
│           prometheus.yml
│           
├───prisma
│   │   schema.prisma
│   │   
│   └───migrations
│       │   migration_lock.toml
│       │   
│       └───20251120033510_init_order_engine
│               migration.sql
│               
├───src
│   │   app.ts
│   │   server.ts
│   │   worker.ts
│   │   
│   ├───config
│   │       bullmq.ts
│   │       db.ts
│   │       env.ts
│   │       redis.ts
│   │       
│   ├───controllers
│   │       order.controler.ts
│   │       
│   ├───queues
│   │       dlq.queue.ts
│   │       order.events.ts
│   │       order.queue.ts
│   │       
│   ├───routes
│   │       ws.route.ts
│   │       
│   ├───services
│   │       dexRouter.ts
│   │       Limitorder.ts
│   │       Sniperorder.ts
│   │       
│   ├───types
│   │       index.ts
│   │       
│   ├───utils
│   │       CircuitBreaker.ts
│   │       logger.ts
│   │       sleep.ts
│   │       
│   ├───worker
│   │       dbsync.worker.ts
│   │       limit.worker.ts
│   │       order.worker.ts
│   │       sniper.worker.ts
│   │       stream.worker.ts
│   │       
│   └───ws
│           websocketManager.ts
│           
└───__tests__
    │   dexRouter.test.ts
    │   orderRoute.test.ts
    │   redisKey.test.ts
    │   setup.ts
    │   teardown.ts
    │   utils.test.ts
    │   websocket.test.ts
    │   
    └───worker
            limitWorker.test.ts
            marketWorker.test.ts
            sniperWorker.test.ts