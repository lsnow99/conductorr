# Profiles

For processing releases, Conductorr has a robust system for defining **profiles**. Profiles are assigned to movies or TV series by the user within the application. Profiles are a way to define exactly what type of release we want to use do download content. 

Profiles consist of both a filtering and a sorting mechanism. The filtering step is used to rule out releases that should not be downloaded under any circumstances. The sorting step is used to decide the order of preference for releases, so that Conductorr can automatically queue the best releases for downloading, and attempt them in order until one succeeds.

To define a profile, users must create both a filter and a sorter. These are defined using a custom scripting language called [CSL](/csl/reference.html). CSL is specifically designed for processing releases has built-in functions to make dealing with them easy. 

CSL scripts are composable, meaning they can [import each other](/csl/reference.html#extension-functions) so that if you have sorting or filtering logic that you want shared between your different profiles, you can do so by importing the shared logic as a function. 

To learn how to build your own CSL scripts, refer to the [documentation on CSL](/csl/reference.html).

If you aren't interesting in building your own scripts, you can use the importing feature to simply invoke scripts that are maintained by the community.
