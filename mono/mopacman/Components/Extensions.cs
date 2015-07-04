using Microsoft.Xna.Framework;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    public static class Extensions
    {
        public static Vector2 ToVector2( this Point p) 
        {
            return new Vector2(p.X, p.Y);
        }

        public static Point ToPoint(this Vector2 p)
        {
            return new Point((int)p.X, (int)p.Y);
        }
    }
}
