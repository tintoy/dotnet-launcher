# A launcher for .NET Core applications (FDD)

Just rename the executable to have the same name as the target assembly (without the `.dll`).
Works regardless of whether the executable is in your `PATH` or symlinked into somewhere on your `PATH`.

If you have questions, or would like to contribute, please feel free to [get in touch](https://github.com/tintoy/dotnet-launcher/issues/new).

-
Note - you could instead run `dotnet publish -r <SomeRID>` to create a [self-contained deployment](https://docs.microsoft.com/en-us/dotnet/articles/core/deploying/index#self-contained-deployments-scd) (which includes a native launcher), but if you want to produce a single directory with executables for more than OS then this might be a better option.

