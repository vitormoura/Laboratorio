using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;
using mopacman.Components;
using mopacman.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Animations
{
    class DirectionalAnimation<T>
        where T : Sprite
    {
        public event EventHandler Finished;
                
        public float Velocity {
            get { return this.velocity; }
            set { this.velocity = Math.Abs(value); }
        }

        public DirectionalAnimation(T c, float velocity)
        {
            this.component = c;
            this.Velocity = velocity;
        }

        public void Start(EnumDirections d, float distance)
        {
            this.remaining = distance;

            switch (d)
            {
                case EnumDirections.North:
                    this.direction = new Vector2(0.0f, distance * -1.0f);
                    break;

                case EnumDirections.South:
                    this.direction = new Vector2(0.0f, distance);
                    break;

                case EnumDirections.East:
                    this.direction = new Vector2(distance, 0.0f);
                    break;

                case EnumDirections.West:
                    this.direction = new Vector2(distance * -1.0f, 0.0f);
                    break;

                default:
                    break;
            }
        }

        public void Update(GameTime gameTime)
        {
            if (this.remaining > 0.0f)
            {
                var elapsed = (float)gameTime.ElapsedGameTime.TotalSeconds;
                var position = this.component.Position;

                float x = (this.direction.X * elapsed * this.velocity);
                float y = (this.direction.Y * elapsed * this.velocity);

                this.remaining -= Math.Abs((this.direction.X != 0) ? x : y);

                if (remaining > 0)
                    this.component.SetPosition(position.X + x, position.Y + y);
            }
            else
            {
                this.remaining = 0.0f;

                if (this.Finished != null)
                    this.Finished.Invoke(this, null);
            }
        }

        private float velocity;
        private float remaining;
        private Vector2 direction;
        private T component;
    }
}
