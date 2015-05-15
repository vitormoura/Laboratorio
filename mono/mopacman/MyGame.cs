﻿using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;
using Microsoft.Xna.Framework.Input;
using mopacman.Components;
using mopacman.Controllers;
using mopacman.Scenes;

namespace mopacman
{
    /// <summary>
    /// This is the main type for your game.
    /// </summary>
    public class MyGame : Game
    {
        public static GraphicsDeviceManager Graphics;
        public static Camera Camera;
        public static SpriteBatch SpriteBatch;

        MazeScene currentScene;
        

        public MyGame()
        {
            Graphics = new GraphicsDeviceManager(this);
            Content.RootDirectory = "Content";
        }

        /// <summary>
        /// Allows the game to perform any initialization it needs to before starting to run.
        /// This is where it can query for any required services and load any non-graphic
        /// related content.  Calling base.Initialize will enumerate through any components
        /// and initialize them as well.
        /// </summary>
        protected override void Initialize()
        {
            // TODO: Add your initialization logic here
            MazeScene scene = new MazeScene(this);
            this.currentScene = scene;
                        
            this.Components.Add(scene);
            
            base.Initialize();
        }

        /// <summary>
        /// LoadContent will be called once per game and is the place to load
        /// all of your content.
        /// </summary>
        protected override void LoadContent()
        {
            MyGame.Graphics.PreferredBackBufferWidth = 465;
            MyGame.Graphics.PreferredBackBufferHeight = 550;
            MyGame.Graphics.IsFullScreen = false;
            MyGame.Graphics.ApplyChanges();

            MyGame.Camera = new Camera(new Vector2(10, 50), MyGame.Graphics.PreferredBackBufferWidth);

            // Create a new SpriteBatch, which can be used to draw textures.
            MyGame.SpriteBatch = new SpriteBatch(GraphicsDevice);
        }

        /// <summary>
        /// UnloadContent will be called once per game and is the place to unload
        /// game-specific content.
        /// </summary>
        protected override void UnloadContent()
        {
            // TODO: Unload any non ContentManager content here
        }

        /// <summary>
        /// Allows the game to run logic such as updating the world,
        /// checking for collisions, gathering input, and playing audio.
        /// </summary>
        /// <param name="gameTime">Provides a snapshot of timing values.</param>
        protected override void Update(GameTime gameTime)
        {
            if (GamePad.GetState(PlayerIndex.One).Buttons.Back == ButtonState.Pressed || Keyboard.GetState().IsKeyDown(Keys.Escape))
                Exit();

            if (Keyboard.GetState().IsKeyDown(Keys.LeftAlt) && Keyboard.GetState().IsKeyDown(Keys.Enter))
            {
                MyGame.Graphics.IsFullScreen = !MyGame.Graphics.IsFullScreen;
                MyGame.Graphics.ApplyChanges();
            }

            elapsed += gameTime.ElapsedGameTime.TotalSeconds;

            this.currentScene.Keyboard.Update(gameTime);

            if (elapsed >= (0.15))
            {
                // TODO: Add your update logic here

                base.Update(gameTime);
                elapsed = 0.0;
            }
            
        }

        /// <summary>
        /// This is called when the game should draw itself.
        /// </summary>
        /// <param name="gameTime">Provides a snapshot of timing values.</param>
        protected override void Draw(GameTime gameTime)
        {
            GraphicsDevice.Clear(Color.Black);

            MyGame.SpriteBatch.Begin();
            base.Draw(gameTime);
            MyGame.SpriteBatch.End();
        }

        private double elapsed;
    }
}
