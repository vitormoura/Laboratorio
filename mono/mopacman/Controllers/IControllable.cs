using mopacman.Components;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Controllers
{
    interface IControllable
    {
        MazeSection CurrentLocation { get; set; }

        MazeSection PreviousLocation { get; }

        void GoTo(EnumDirections d);
    }
}
