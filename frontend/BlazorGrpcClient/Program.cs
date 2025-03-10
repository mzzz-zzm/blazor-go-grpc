using BlazorGrpcClient;
using Grpc.Net.Client;
using Grpc.Net.Client.Web;
using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Web;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;

var builder = WebAssemblyHostBuilder.CreateDefault(args);
builder.RootComponents.Add<App>("#app");
builder.RootComponents.Add<HeadOutlet>("head::after");

// Add logging
builder.Logging.SetMinimumLevel(LogLevel.Debug);

// Add regular HttpClient
builder.Services.AddScoped(sp => new HttpClient { BaseAddress = new Uri(builder.HostEnvironment.BaseAddress) });

// Configure the gRPC-Web client
builder.Services.AddSingleton(services =>
{
    // Get the server URL - adjust based on environment
    // var serverUrl = builder.HostEnvironment.IsDevelopment()
    //     ? "http://localhost:8080" 
    //     : "http://backend:8080";

    var serverUrl = "http://localhost:8080"; // TODO:MH builder.HostEnvironment.IsDevelopment() doesn't pickup env var... bug?
        
    Console.WriteLine($"-----gRPC Server URL: {serverUrl} host: {builder.HostEnvironment.BaseAddress}");
    
    // Create a gRPC-Web handler
    var httpHandler = new GrpcWebHandler(GrpcWebMode.GrpcWeb, new HttpClientHandler());
    
    // Create and return the gRPC channel
    return GrpcChannel.ForAddress(serverUrl, new GrpcChannelOptions
    {
        HttpHandler = httpHandler,
        // For debugging
        LoggerFactory = services.GetRequiredService<ILoggerFactory>()
    });
});

// Register the Greeter client
builder.Services.AddSingleton(services => 
    new BlazorGrpcClient.proto.Greeter.GreeterClient(services.GetRequiredService<GrpcChannel>()));

await builder.Build().RunAsync();